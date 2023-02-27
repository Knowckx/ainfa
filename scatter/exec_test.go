package scatter

import (
	"bytes"
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
	in := "sh test.sh"
	// in := "kubectl -n smec get secret eureka -o yaml"
	cmd := NewCommand(in)
	var outb, errb bytes.Buffer
	cmd.Stderr = &errb
	cmd.Stdout = &outb
	err := cmd.RunCommand() // 这个错误没收集全
	fmt.Println(err)
	fmt.Println("stdout:", outb.String(), "stderr:", errb.String())
}

func Test_CompressSpaces(t *testing.T) {
	in := "  	l  s  		-l	"
	out := util.CompressSpaces(in)
	fmt.Println(out)
}
