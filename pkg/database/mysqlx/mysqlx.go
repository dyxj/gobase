package mysqlx

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var (
	mainDb      *sqlx.DB
	mainConfig  Config
	mainOptions []Option
)

// Config contains mandatory fields to establish a connection
//
// note: not convinced that interfaces are the right approach
// could use struct that has validate method instead
type Config interface {
	Logger() *logrus.Logger
	MySqlDsn() string
}

// Option manipulates default values of DB
type Option func(db *sqlx.DB)

// NewDb initializes database connections.
// Should be called in main
var NewDb = func(cfg Config, options ...Option) *sqlx.DB {
	mainConfig = cfg
	mainOptions = options
	return GetDb()
}

// GetDb can be used if this package is included in dao layer
//
// note: should revert to original method, don't like the idea of panics here
var GetDb = func() *sqlx.DB {
	if mainDb == nil {
		logger().Infof("initializing database")
		var err error
		mainDb, err = sqlx.Open("mysql", mainConfig.MySqlDsn())
		if err != nil {
			logger().Panicf("failed to open database connection: %v", err)
		}

		err = mainDb.Ping()
		if err != nil {
			logger().Panicf("failed to ping database: %v", err)
		}

		for _, opt := range mainOptions {
			opt(mainDb)
		}
		logger().Infof("database connected successfully")
	}
	return mainDb
}

// CloseDb should be called with defer after NewDb()
var CloseDb = func() {
	logger().Infof("closing database")
	if mainDb != nil {
		err := mainDb.Close()
		if err != nil {
			logger().Errorf("error closing database: %v", err)
		}
	}
	logger().Infof("database closed successfully")
}

// note: New type could be better here to tag logs
func logger() *logrus.Entry {
	return mainConfig.Logger().WithFields(logrus.Fields{
		"pkg": "mysqlx",
	})
}
