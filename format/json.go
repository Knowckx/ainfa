package format

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rs/zerolog/log"
)

// save json | load json
func LoadJsonFile(f string, data interface{}) error {
	out, err := ioutil.ReadFile(f)
	if err != nil {
		log.Error().Str("file", f).Err(err).Msg("read file error")
		return err
	}
	err = json.Unmarshal(out, data)
	if err != nil {
		log.Error().Str("file", f).Err(err).Msg("yaml.Unmarshal error")
		return err
	}
	return err
}
