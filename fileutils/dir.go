package fileutils

import(
	"os"
	"strings"
	"path/filepath"
	"fmt"
)


// Expand ~ to value of ENV HOME
func expandTilde(f string) string {
	if strings.HasPrefix(f, "~"+string(filepath.Separator)) {
		return os.Getenv("HOME") + f[1:]
	}
	return f
}

// Check for path and create if it doesn't exist
func CreatePath(path string) (string, error) {
	absPath,err := filepath.Abs(expandTilde(path))
	if err != nil {
		return "", err
	}
	stat,err := os.Stat(absPath) 
	if stat == nil && os.IsNotExist(err) {
		err := os.MkdirAll(absPath, 0755)
		if err != nil {
			return "", err
		}
	} else if !stat.IsDir() {
		err := fmt.Errorf("path: %#v exists but is not a directory\n", path)
		return "", err
	}

	return absPath, nil
}

