package scatter

import (
	"fmt"
	"testing"

	"github.com/Knowckx/ainfa/util"
)

func Test_ParseExec(t *testing.T) {
	in := "  ls  -l "
	out := ParseExec(in)
	fmt.Println(out)
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
	out := util.CompressSpaces(in)
	fmt.Println(out)
}
