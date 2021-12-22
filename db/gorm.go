package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GenDefaultConfig() *gorm.Config {
	logConfig := logger.Config{
		SlowThreshold: time.Second,  // Slow SQL threshold
		LogLevel:      logger.Error, // Log level
		Colorful:      false,        // Disable color
	}
	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logConfig,
	)
	out := &gorm.Config{
		Logger:                 log,
		SkipDefaultTransaction: true,
	}
	return out
}

type PGConnInfo struct {
	HOST     string
	PORT     string
	DATABASE string
	USER     string
	PASWD    string
}

func (in PGConnInfo) String() string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable  application_name=%s TimeZone=Asia/Shanghai",
		in.HOST, in.PORT, in.USER, in.PASWD, in.DATABASE, "gorm",
	)
	return dsn
}

func ConnPG(in *PGConnInfo) (db *gorm.DB, err error) {
	dsn := in.String()
	d := postgres.Open(dsn)
	pgConfig := GenDefaultConfig()
	db, err = gorm.Open(d, pgConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
