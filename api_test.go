package infa

import (
	"fmt"
	"testing"
)

func Test_LocFilePath(t *testing.T) {
	out := LocFilePath("ainfa", "util", "util.go")
	fmt.Println(out)
}
