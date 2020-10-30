package path

import (
	"os"
	"path/filepath"
	"strings"
)

func ResolveModulePath() (string, error) {
	execFilePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	modulePathParts := strings.Split(filepath.Dir(execFilePath), "/")
	return strings.Join(modulePathParts[0:len(modulePathParts)-1], "/"), nil
}
