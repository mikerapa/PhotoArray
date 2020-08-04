package PhotoArray

import (
	"os"
	"strings"
)

func IsValidFilePath(path string) bool {
	path = strings.Trim(path, " ")
	if len(path) == 0 {
		return false
	}

	fileItem, err := os.Stat(path)
	return !os.IsNotExist(err) && fileItem.Mode().IsRegular()
}
