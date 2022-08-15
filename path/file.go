package path

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func SaveDataToFile(fileName, data string) error {
	path := fileName // "./" +
	f, err := os.Create(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	f.WriteString(data)
	log.Info().Str("fileName", fileName).Msg("write data to file success.")
	return nil
}

func ReadFile(f string) (string, error) {
	res, err := ioutil.ReadFile(f)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(res), nil
}
