package golang

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/coding-wepack/wepack-cli/pkg/log"
	"github.com/coding-wepack/wepack-cli/pkg/log/logfields"
	"github.com/coding-wepack/wepack-cli/pkg/settings"
	"github.com/coding-wepack/wepack-cli/pkg/util/httputil"
	"github.com/coding-wepack/wepack-cli/pkg/util/ioutils"
	"github.com/coding-wepack/wepack-cli/pkg/util/sliceutil"
	"github.com/coding-wepack/wepack-cli/pkg/util/sysutil"
	"github.com/pkg/errors"
	"golang.org/x/mod/semver"
)

var excludeList = []string{".idea", ".git", "vendor"}

func Push() error {
	log.Infof("begin to publish go artifacts %s to %s", settings.Module, settings.Repo)
	if settings.Verbose {
		log.Debug("registry info",
			logfields.String("repo", settings.Repo),
			logfields.String("username", settings.Username),
			logfields.String("password", settings.Password))
	}

	// 解析 mod.path 和 mod.version
	modPath, modVer, err := parseModule()
	if err != nil {
		return err
	}
	if settings.Verbose {
		log.Debugf("mod.path: %s, mod.version: %s", modPath, modVer)
	}

	log.Info("check go.mod file is in the root dir")
	// 获取项目下的文件列表
	filePaths, err := findFileList("./", excludeList)
	if err != nil {
		return err
	}

	// 校验当前目录中是否存在 go.mod 文件
	if !sliceutil.ContainsInStringSlice(filePaths, "go.mod") {
		return errors.New("the go.mod file must be included in the root dir")
	}

	log.Infof("processing file path and making zip package...")
	// 此处使用代码创建 zip 包可以摆脱对执行环境的差异以及是否安装 zip 的影响
	// 创建一个 zip 压缩文件
	zipFile, err := os.CreateTemp("./", "*.zip")
	if err != nil {
		return err
	}
	defer func() {
		_ = zipFile.Close()
		_ = os.Remove(zipFile.Name())
	}()
	// 创建一个 zip.Writer
	zipWriter := zip.NewWriter(zipFile)
	defer func() { _ = zipWriter.Close() }()

	for _, path := range filePaths {
		// 将文件写入
		err = writeFile(path, zipWriter)
		if err != nil {
			return err
		}
	}
	_ = zipWriter.Close()
	_, _ = zipFile.Seek(0, 0)

	// 上传文件
	url := fmt.Sprintf("%s/%s/@v/%s.zip", strings.Trim(settings.Repo, "/"), modPath, modVer)
	log.Info("analyzing the zip package and uploading it to the remote repository")
	if settings.Verbose {
		log.Debugf("upload to remote repo url: %s", url)
	}
	resp, err := httputil.DefaultClient.Put(url, "", zipFile, settings.Username, settings.Password)
	if err != nil {
		return err
	}
	defer ioutils.QuiteClose(resp.Body)
	if resp.StatusCode >= http.StatusBadRequest {
		return errors.Errorf("got an unexpected response status: %s", resp.Status)
	}
	log.Infof("artifacts upload successful!")
	return nil
}

func writeFile(filePath string, zipWriter *zip.Writer) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Error("file close failed", logfields.Error(err))
		}
	}()
	// 创建一个新的文件头并设置前缀
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}
	header.Name = getZipFileHeaderName(filePath)
	if settings.Verbose {
		log.Debugf("zip file %s", header.Name)
	}
	// 将文件头写入压缩包
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	// 将文件内容写入压缩包
	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}
	return nil
}

// parseModule 用于从 module 中解析出 mod.path 和 mod.version
func parseModule() (modPath, version string, err error) {
	split := strings.Split(settings.Module, "@")
	if len(split) != 2 {
		err = errors.New("invalid module must conform to the {mod.path}@{version} format, eg: coding.net/wepack/cli@v0.0.1")
		return
	}
	modPath, version = split[0], split[1]
	if !semver.IsValid(version) {
		err = errors.New("unrecognized version")
	}
	return
}

// findFileList 用于获取指定路径下的文件列表
func findFileList(dir string, excludeList []string) (filePaths []string, err error) {
	needFilter := len(excludeList) > 0
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			if needFilter {
				for _, e := range excludeList {
					if strings.EqualFold(path, e) || strings.HasPrefix(path, e) {
						return err
					}
				}
			}
			filePaths = append(filePaths, path)
		}
		return nil
	})
	return
}

func getZipFileHeaderName(filePath string) string {
	if sysutil.IsWindows() {
		filePath = strings.ReplaceAll(filePath, "\\", "/")
	}
	return fmt.Sprintf("%s/%s", settings.Module, filePath)
}
