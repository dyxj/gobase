package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

const (
	// Environment variable prefix
	evPrefix = ""
)

// FromEnvVar creates Config from environment variables.
// Defaults are not set and all environment variables are marked as required.
// Sample, recommended or valid values are provided in "desc" tag.
func FromEnvVar() (*Config, error) {
	var cfg Config
	err := envconfig.Process(evPrefix, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}

// Help prints to std out environment variable configuration
func Help() error {
	var cfg Config
	return envconfig.Usage(evPrefix, &cfg)
}

// Config stores service configurations
//
// note: create new struct and inject logger
// other methods to satisfy config interfaces should also be tied to new struct
type Config struct {
	// Variables
	Svc service
	Log log
	Db  database

	// Logger
	initLogger sync.Once
	logger     *logrus.Logger
}

// Logger returns singleton logger
func (c *Config) Logger() *logrus.Logger {
	c.initLogger.Do(func() {
		c.logger = logrus.New()
		c.logger.SetLevel(c.Log.Level)
	})
	return c.logger
}

// MySqlDsn generates mysql dsn from config variables
func (c *Config) MySqlDsn() string {
	opt := "&parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=%s",
		c.Db.User, c.Db.Pass, c.Db.Host, c.Db.Port, c.Db.Name, c.Db.SSL) + opt
	return dsn
}

type service struct {
	Profile string        `required:"true" desc:"[local,prod] service profile "`
	Port    string        `required:"true" desc:"[8080] server port "`
	Timeout time.Duration `required:"true" desc:"[2m] server timeout "`
}

type log struct {
	Level logrus.Level `required:"true" desc:"[info,debug] log levels "`
}

type database struct {
	User            string        `required:"true" desc:"[mysqluser]"`
	Pass            string        `required:"true" desc:"[abc123]"`
	Name            string        `required:"true" desc:"[gobase]"`
	Host            string        `required:"true" desc:"[127.0.0.1]"`
	Port            string        `required:"true" desc:"[3306]"`
	SSL             string        `required:"true" desc:"[true,false,skip-verify,preferred]"`
	MaxOpenConn     int           `required:"true" desc:"[20]"`
	MaxIdleConn     int           `required:"true" desc:"[5]"`
	MaxConnLifetime time.Duration `required:"true" desc:"[50m]"`
}
