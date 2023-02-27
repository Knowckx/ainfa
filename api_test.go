package infa

import (
	"fmt"
	"testing"
)

func Test_LocFilePath(t *testing.T) {
	out := LocFilePath("in-fa", "util", "util.go")
	fmt.Println(out)
}
