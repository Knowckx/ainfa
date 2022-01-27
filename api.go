package infa

import (
	"github.tools.sap/aeolia/in-fa/path"
	"github.tools.sap/aeolia/in-fa/scatter"
)

func SaveDataToFile(fileName, data string) error {
	return path.SaveDataToFile(fileName, data)
}

func ExecCmd(in string) (string, error) {
	return scatter.ExecCmd(in)
}
