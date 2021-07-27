package infa

import (
	"os"
	"path/filepath"
	"strings"
)

func GetRootPathProjName(projName string) string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if projName == "" || !strings.Contains(path, projName) {
		return path
	}
	paths := strings.Split(path, projName)
	execPath := filepath.Join(paths[0], projName)
	return execPath
}

func GetRootPath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}
