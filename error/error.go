package error

import (
	"fmt"
	"os"
)

const (
	LF = 10
)

func ErrorAndExit(format string, args ...interface{}) {
	if format[0] == LF {
		fmt.Fprintf(os.Stderr, format[1:]+"\n", args...)
	} else {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("ERROR: %s\n", format), args...)
	}
	os.Exit(1)
}

func ExitIfError(err error, format string, args ...interface{}) {
	if err != nil {
		ErrorAndExit(format, args)
	}
}

func ExitIfNoSourceDir(err error) {
	if err != nil {
		if os.IsNotExist(err) {
			ErrorAndExit("Current directory does not contain a metadata or src directory")
		}

		ErrorAndExit(err.Error())
	}
}
