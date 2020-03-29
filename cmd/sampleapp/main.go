package main

import (
	"github.com/dyxj/gobase/config"
	"github.com/dyxj/gobase/pkg/database/mysqlx"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func main() {
	err := config.Help()
	if err != nil {
		logrus.Fatalf("broken build: %v", err)
	}

	cfg, err := config.FromEnvVar()
	if err != nil {
		logrus.Fatalf("config.FromEnvVar: %v", err)
	}

	db := mysqlx.NewDb(cfg, func(db *sqlx.DB) {
		db.SetMaxOpenConns(cfg.Db.MaxOpenConn)
		db.SetMaxIdleConns(cfg.Db.MaxIdleConn)
		db.SetConnMaxLifetime(cfg.Db.MaxConnLifetime)
	})
	defer mysqlx.CloseDb()

	// to compile
	cfg.Logger().Printf("%+v", db)
}
