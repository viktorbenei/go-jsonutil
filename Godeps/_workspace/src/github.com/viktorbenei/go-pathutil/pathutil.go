package pathutil

import (
	"os"
	"strings"
)

func IsRelativePath(pth string) bool {
	if strings.HasPrefix(pth, "./") {
		return true
	}

	if strings.HasPrefix(pth, "/") {
		return false
	}

	if strings.HasPrefix(pth, "$") {
		return false
	}

	return true
}

func IsPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
