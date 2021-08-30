package path

import (
	"os"
	"path/filepath"
	"strings"
)

func GetProjRootPath(projName string) string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if projName == "" || !strings.Contains(path, projName) {
		return path
	}
	paths := strings.Split(path, projName)
	execPath := filepath.Join(paths[0], projName)
	execPath = execPath + "/"
	return execPath
}

func GetProjFilePath(mids ...string) string {
	if len(mids) < 2 {
		return ""
	}
	root := GetProjRootPath(mids[0])
	rest := filepath.Join(mids[1:]...)
	out := filepath.Join(root, rest)
	out = out + "/"
	return out
}

func GetRootPath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}
