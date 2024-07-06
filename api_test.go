package infa

import (
	"fmt"
	"testing"
)

func Test_LocFilePath(t *testing.T) {
	out := LocFilePath("in-fa", "util", "util.go")
	fmt.Println(out)
}

func Test_V3(t *testing.T) {
	a := Return(1 > 1, 1, 2)
	fmt.Println(a)
}
