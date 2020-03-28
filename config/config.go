package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

const (
	// Environment variable prefix
	evPrefix = ""
)

// Config stores service configurations
type Config struct {
	Svc service
	Log log
	Db  database
}

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

type service struct {
	Profile string        `required:"true" desc:"[local,prod] service profile "`
	Port    string        `required:"true" desc:"[8080] server port "`
	Timeout time.Duration `required:"true" desc:"[2m] server timeout "`
}

type log struct {
	Level string `required:"true" desc:"[info,debug] log levels "`
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
