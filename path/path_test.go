package path

import (
	"fmt"
	"testing"
)

func Test_GetProjRootPath(t *testing.T) {
	res := GetProjRootPath("in-fa")
	fmt.Println("result:", res)
	return
}

func Test_GetProjFilePath(t *testing.T) {
	res := GetProjFilePath("in-fa", "config")
	fmt.Println(res)
	return
}

func Test_GetRootPath(t *testing.T) {
	res := GetRootPath()
	fmt.Println(res)
	return
}
