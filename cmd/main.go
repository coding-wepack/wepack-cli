package main

import (
	"fmt"
	"os"

	"github.com/coding-wepack/wepack-cli/pkg/log"
	"github.com/coding-wepack/wepack-cli/pkg/settings"
)

func main() {
	defer log.Sync()

	cmd, err := newRootCmd()
	if err != nil {
		log.Warn(err.Error())
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		// debug("%+v", err)
		os.Exit(1)
	}
}

func info(format string, v ...interface{}) {
	format = fmt.Sprintf("%s\n", format)
	_, _ = fmt.Fprintf(os.Stdout, format, v...)
}

func debug(format string, v ...interface{}) {
	if settings.Verbose {
		format = fmt.Sprintf("%s", format)
		log.Debug(fmt.Sprintf(format, v...))
	}
}

func warning(format string, v ...interface{}) {
	format = fmt.Sprintf("WARNING: %s\n", format)
	_, _ = fmt.Fprintf(os.Stderr, format, v...)
}
