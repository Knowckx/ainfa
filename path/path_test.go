package path

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_GetProjRootPath(t *testing.T) {
	res := GetProjRootPath("in-fa")
	fmt.Println(res)
	return
}

func Test_GetFilePath(t *testing.T) {
	res := GetFilePath("testfile", "azure.toml")
	fmt.Println(res)
	return
}

type AuthParams struct {
	Host string
	User string
}

func GenTestFilesPath(fileName string) string {
	rootPath := GetRootPath()
	out := filepath.Join(rootPath, "testfiles", fileName)
	return out
}
