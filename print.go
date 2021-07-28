package infa

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pelletier/go-toml/v2"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

func Printf(format string, a ...interface{}) (n int, err error) {
	format = fmt.Sprintf("%s\n", format)
	return fmt.Printf(format, a...)
}

func PrintJson(in interface{}) {
	res, err := json.Marshal(in)
	if err != nil {
		log.Error().Stack().Err(err).Send()
		return
	}
	Printf(string(res))
}

func PrintYaml(in interface{}) {
	res, err := yaml.Marshal(in)
	if err != nil {
		log.Error().Stack().Err(err).Send()
		return
	}
	Printf(string(res))
}

func PrintToml(in interface{}) {
	res, err := toml.Marshal(in)
	if err != nil {
		log.Error().Stack().Err(err).Send()
		return
	}
	Printf(string(res))
}

func Print(x interface{}) {
	Printf("Input Type:%T", x)
	displayPath("value", reflect.ValueOf(x))
}
