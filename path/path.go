package path

import (
	"os"
	"path/filepath"
	"strings"
)

func GetAppPath() string {
	dir, _ := os.Getwd()

	exePath := os.Args[0]
	if len(exePath) == 0 || strings.HasSuffix(exePath, ".test") {
		return dir
	}

	sep := string(os.PathSeparator)
	sepCount := strings.Count(exePath, sep)
	if sepCount <= 1 {
		return dir
	}

	idx1 := strings.Index(exePath, sep)
	idx2 := strings.LastIndex(exePath, sep)

	add := exePath[idx1:idx2]
	out := dir + add
	return out
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
