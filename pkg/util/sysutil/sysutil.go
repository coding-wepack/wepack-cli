package sysutil

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

const (
	LINUX   = "linux"
	WINDOWS = "windows"
)

func GoosType() (string, error) {
	sysType := runtime.GOOS

	if strings.EqualFold(sysType, LINUX) {
		return LINUX, nil
	} else if strings.EqualFold(sysType, WINDOWS) {
		return WINDOWS, nil
	} else {
		return "", errors.New(fmt.Sprintf("unknown GOOS type: %s", sysType))
	}
}

func IsWindows() bool {
	return strings.EqualFold(runtime.GOOS, WINDOWS)
}

func IsLinux() bool {
	return strings.EqualFold(runtime.GOOS, LINUX)
}
