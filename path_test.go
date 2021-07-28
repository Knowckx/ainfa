package infa

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_GetProjRootPath(t *testing.T) {
	res := GetRootPathProjName("azuretest")
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

func Test_LoadYamlFile(t *testing.T) {
	out := &AuthParams{}
	fPath := GenTestFilesPath("userinfo.yaml")
	err := LoadYamlFile(fPath, out)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	return
}

func Test_LoadTomlFile(t *testing.T) {
	out := map[string]AuthParams{}
	fPath := GenTestFilesPath("azure.toml")
	err := LoadTomlFile(fPath, &out)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	return
}
