package infa

import (
	"fmt"
	"testing"
)

func Test_LocFilePath(t *testing.T) {
	out := LocFilePath("in-fa", "util", "util.go")
	fmt.Println(out)
}

func Test_QuickTest(t *testing.T) {
	avger := Averager{}
	avger.AddNumber(1)
	avger.AddNumber(2.1)
	avger.AddNumber(5)
	fmt.Println(avger.Avg)
}
