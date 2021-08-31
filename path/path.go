package path

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetAppPath() string {
	fmt.Println(os.Args)
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
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
