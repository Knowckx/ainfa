package infa

import (
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
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
	cmd := CompressSpaces(in)
	cmd = strings.TrimSpace(cmd)
	return cmd
}

func ParseToFragments(in string) []string {
	return strings.Split(in, " ")
}

// Compress Spaces and Tap to one space
func CompressSpaces(in string) string {
	out := ""
	var added int
	for _, key := range in {
		if key == ' ' || key == '	' {
			if added == 0 {
				out += " "
				added = 1
			} else {
				added++
			}
		} else {
			out += string(key)
			added = 0
		}
	}
	return out
}
