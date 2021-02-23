package config

import "github.com/kelseyhightower/envconfig"

type OsEnvironment struct {
	Env              string `default:"local"`
	MysqlUser        string `default:"pantogram" split_words:"true"`
	MysqlPassword    string `default:"pantogram" split_words:"true"`
	MysqlDbname      string `default:"pantogram" split_words:"true"`
	MysqlUrl         string `default:"127.0.0.1:53306" split_words:"true"`
	InfluxdbUser     string `default:"pantogram" split_words:"true"`
	InfluxdbPassword string `default:"pantogram" split_words:"true"`
	InfluxdbDbname   string `default:"pantogram" split_words:"true"`
	InfluxdbUrl      string `default:"127.0.0.1:8086" split_words:"true"`
}

var Environment OsEnvironment

func init() {
	envconfig.Process("", &Environment)
}
