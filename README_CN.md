## wp
`wp` 是一个帮助您管理 [WePack](https://wepack.coding.net) 制品的客户端工具，它的全称是：`WePack-CLI`。

例如 `go` 命令，使用 `wp go` 可以用于推送本地 go 制品到远程仓库，或者拉取远程仓库的 go 制品到本地。

目前 `wp` 支持管理 `go` 语言制品，支持您将本地 go 代码推送到 WePack 的制品仓库中


## 安装 wp

### cURL & wget
```shell
# Linux 
$ curl -fL "https://coding-public-generic.pkg.coding.net/registry/disk/wp/linux/amd64/wp?version=latest" -o wp
# 或者使用 wget 方式
$ wget "https://coding-public-generic.pkg.coding.net/registry/disk/wp/linux/amd64/wp?version=latest" -O wp

# MacOS
$ curl -fL "https://coding-public-generic.pkg.coding.net/registry/disk/wp/macos/amd64/wp?version=latest" -o wp
# 或者使用 wget 方式
$ wget "https://coding-public-generic.pkg.coding.net/registry/disk/wp/macos/amd64/wp?version=latest" -O wp

# Windows
$ curl -fL "https://coding-public-generic.pkg.coding.net/registry/disk/wp/darwin/amd64/wp?version=latest" -o wp
# 或者使用 wget 方式
$ wget "https://coding-public-generic.pkg.coding.net/registry/disk/wp/darwin/amd64/wp?version=latest" -O wp

$ chmod +x wp
$ sudo mv wp /usr/local/bin

# 检验是否成功安装
$ wp
```

## 帮助文档
```shell
$ wp

The WePack Artifacts Manager Client

The migrate argument must be an artifact type, available now:
- go

Common actions for wp:
- wp go: Manager go artifacts

Usage:
  wp [command]

Available Commands:
  go          manager go artifacts.
  help        Help about any command
  version     Print the CLI version

Flags:
  -h, --help      help for wp
  -v, --verbose   Make the operation more talkative

Use "wp [command] --help" for more information about a command.
```

## 制品

### go
```shell
# 推送 go 制品到 WePack 制品仓库
$ wp go push --module github.com/coding-wepack/wepack-cli@v0.0.1 --repo "https://demo-go.pkg.wepack.net/project/repo/" -u test -p test_pwd
2023-05-26 15:45:15.678 INFO    begin to publish go artifacts github.com/coding-wepack/wepack-cli@v0.0.1 to https://demo-go.pkg.wepack.net/project/repo/
2023-05-26 15:45:15.678 INFO    check go.mod file is in the root dir
2023-05-26 15:45:15.742 INFO    processing file path and making zip package...
2023-05-26 15:45:15.758 INFO    analyzing the zip package and uploading it to the remote repository
2023-05-26 15:45:18.202 INFO    artifacts upload successful!

$ wp go push help
This command publish artifacts to a WePack Artifact Registry.

Examples:
    # Publish go artifacts to a WePack Artifact Registry:
    $ wp go push --module github.com/coding-wepack/wepack-cli@v0.0.1 --repo "https://demo-go.pkg.wepack.net/project/repo/" -u test -p test_pwd

Flags '--module' and '--repo' must be set.

Usage:
  wp go push [flags]

Flags:
  -h, --help              help for push
  -m, --module string     e.g., --module github.com/coding-wepack/wepack-cli@v0.0.1 or -m github.com/coding-wepack/wepack-cli@v0.0.1
  -p, --password string   e.g., --password test_pwd or -u test_pwd
  -r, --repo string       e.g., --repo https://demo-go.pkg.wepack.net/project/repo/ or -r https://demo-go.pkg.wepack.net/project/repo/
  -u, --username string   e.g., --username test or -u test

Global Flags:
  -v, --verbose   Make the operation more talkative
```
