package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// resolveOutputPath determines where to save the file
func ResolveOutputPath(filePath, destPath string) (string, error) {
	filename := filepath.Base(filePath)

	if destPath == "" {
		return filename, nil
	}

	if strings.HasSuffix(destPath, "/") {
		if _, err := os.Stat(destPath); os.IsNotExist(err) {
			return "", fmt.Errorf("directory %s does not exist", destPath)
		}
		return filepath.Join(destPath, filename), nil
	}

	dir := filepath.Dir(destPath)
	if dir != "." {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return "", fmt.Errorf("directory %s does not exist", dir)
		}
	}
	return destPath, nil
}
