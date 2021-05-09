package infrastructure

import (
	"github.com/kelseyhightower/envconfig"
)

type osEnvironment struct {
	Env           string `default:"local"`
	MysqlUser     string `default:"pantogram" split_words:"true"`
	MysqlPassword string `default:"pantogram" split_words:"true"`
	MysqlDbname   string `default:"pantogram" split_words:"true"`
	MysqlUrl      string `default:"127.0.0.1:53306" split_words:"true"`
}

var environment osEnvironment

func init() {
	envconfig.Process("", &environment)
}
