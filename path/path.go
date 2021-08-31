package path

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetAppPath() string {
	_, ss, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(ss)
	return basepath
}

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

func GetFilePath(root string, mids ...string) string {
	rest := filepath.Join(mids...)
	out := filepath.Join(root, rest)
	return out
}
