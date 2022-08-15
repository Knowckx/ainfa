package infa

import (
	"github.com/Knowckx/ainfa/format"
	"github.com/Knowckx/ainfa/parallel"
	"github.com/Knowckx/ainfa/path"
	"github.com/Knowckx/ainfa/scatter"
	"github.com/Knowckx/ainfa/util"
)

func ShortStr(in string) string {
	return util.ShortStr(in)
}

func SaveDataToFile(fileName, data string) error {
	return path.SaveDataToFile(fileName, data)
}

func LoadJsonFile(f string, data interface{}) error {
	return format.LoadJsonFile(f, data)
}

func ReadFile(f string) (string, error) {
	return path.ReadFile(f)
}

func ExecCmd(in string) (string, error) {
	return scatter.ExecCmd(in)
}

func LocFilePath(projName string, mids ...string) string {
	return path.LocFilePath(projName, mids...)
}

func NewParallel(max int) *parallel.Parallel {
	return parallel.NewParallel(max)
}
