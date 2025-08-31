package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostgresqlDB struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	User         string `mapstructure:"user" json:"user" yaml:"user"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Schema       string `mapstructure:"schema" json:"schema" yaml:"schema"`
	TimeZone     string `mapstructure:"time-zone" json:"time-zone" yaml:"time-zone"`
	SslMode      string `mapstructure:"sslmode" json:"sslmode" yaml:"sslmode"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
}

type Zap struct {
	Level     string `mapstructure:"level" json:"level" yaml:"level"`
	Path      string `mapstructure:"path" json:"path" yaml:"path"`
	MaxSize   int    `mapstructure:"max-size" json:"max-size" yaml:"max-size"`
	MaxBackup int    `mapstructure:"max-backup" json:"max-backup" yaml:"max-backup"`
	MaxAge    int    `mapstructure:"max-age" json:"max-age" yaml:"max-age"`
	Compress  bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type Config struct {
	Domain   string       `yaml:"domain" json:"domain" mapstructure:"domain"`
	Postgres PostgresqlDB `yaml:"postgres" json:"postgres" mapstructure:"postgres"`
	Zap      Zap          `yaml:"zap" json:"zap" mapstructure:"zap"`
}

var (
	SELF_CONFIG Config
	SELF_DB     *gorm.DB
	SELF_VIPER  *viper.Viper
	SELF_LOG    *zap.Logger
)
