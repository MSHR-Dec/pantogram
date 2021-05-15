package infrastructure

import "github.com/kelseyhightower/envconfig"

type OsEnvironment struct {
	Env              string `default:"local"`
	GRPCListenPort   string `default:":8081" split_words:"true"`
	InfluxdbUser     string `default:"pantogram" split_words:"true"`
	InfluxdbPassword string `default:"pantogram" split_words:"true"`
	InfluxdbDbname   string `default:"pantogram" split_words:"true"`
	InfluxdbUrl      string `default:"127.0.0.1:8086" split_words:"true"`
}

var Environment OsEnvironment

func init() {
	envconfig.Process("", &Environment)

	if Environment.Env != "local" {
		Environment.GRPCListenPort = ":8080"
	}
}
