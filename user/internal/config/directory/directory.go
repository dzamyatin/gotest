package directory

import (
	"path/filepath"
	"runtime"
)

func VarDir() string {
	return ConfigDir() + "/../../var"
}

func ConfigDir() string {
	return currDir() + "/.."
}

func currDir() string {
	_, file, _, ok := runtime.Caller(0)

	if ok {
		return filepath.Dir(file)
	}

	panic("Can't get config directory")
}
