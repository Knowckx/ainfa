package infa

import (
	"fmt"
	"testing"
)

func Test_ParseExec(t *testing.T) {
	in := "  ls  -l "
	out := ParseExec(in)
	PrintJson(out)
}

func Test_RunCommand(t *testing.T) {
	// in := "  ls  -l "
	// in := "kubecm switch eureka--cd"
	in := "kubectl -n smec get secret eureka -o yaml"
	cmd := NewCommand(in)
	cmd.Run()
}

func Test_CompressSpaces(t *testing.T) {
	in := "  	l  s  		-l	"
	out := CompressSpaces(in)
	fmt.Println(out)
}
