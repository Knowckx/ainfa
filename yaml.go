package infa

import (
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

func LoadYamlFile(f string, data interface{}) error {
	out, err := ioutil.ReadFile(f)
	if err != nil {
		log.Error().Str("file", f).Err(err).Msg("read file error")
		return err
	}
	err = yaml.Unmarshal(out, data)
	if err != nil {
		log.Error().Str("file", f).Err(err).Msg("yaml.Unmarshal error")
		return err
	}
	return err
}
