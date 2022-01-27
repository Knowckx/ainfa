package scatter

import (
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.tools.sap/aeolia/in-fa/util"
)

type Command struct {
	CmdSrc    string
	Fragments []string
	ExecCmd   *exec.Cmd
}

func ExecCmd(in string) (string, error) {
	cmd := NewCommand(in)
	return cmd.Run()
}

func NewCommand(in string) *Command {
	out := &Command{}
	out.CmdSrc = ParseExec(in)
	out.Fragments = ParseToFragments(out.CmdSrc)
	return out
}

func (t *Command) Run() (string, error) {
	cmdLen := len(t.Fragments)
	if cmdLen == 1 {
		t.ExecCmd = exec.Command(t.Fragments[0])
	} else {
		t.ExecCmd = exec.Command(t.Fragments[0], t.Fragments[1:cmdLen]...)
	}
	return t.RunCommand()
}

func (t *Command) RunCommand() (string, error) {
	log.Info().Msgf("Run Command: %s", t.CmdSrc)
	info, err := t.ExecCmd.Output()
	// info, err := t.ExecCmd.CombinedOutput()
	if err != nil {
		return "", errors.WithStack(err)
	}
	res := string(info)
	return res, nil
}

func ParseExec(in string) string {
	cmd := util.CompressSpaces(in)
	cmd = strings.TrimSpace(cmd)
	return cmd
}

func ParseToFragments(in string) []string {
	return strings.Split(in, " ")
}
