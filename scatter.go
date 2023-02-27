package infa

import (
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Warn().Str("key", key).Msg("Get environment variable failed. return nil.")
	}
	return value
}

func FormatTime(in time.Time) string {
	return in.Format(time.RFC3339)
}

// server time use zone Shanghai.
func SHTime(in time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return in.In(loc)
}
