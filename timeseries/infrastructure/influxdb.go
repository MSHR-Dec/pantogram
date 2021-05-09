package infrastructure

import (
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InitInfluxdb() (client influxdb2.Client) {
	option := influxdb2.DefaultOptions()
	if Environment.Env != "prd" {
		option.SetLogLevel(3)
	}

	client = influxdb2.NewClientWithOptions(
		fmt.Sprintf("http://%s", Environment.InfluxdbUrl),
		fmt.Sprintf("%s:%s", Environment.InfluxdbUser, Environment.InfluxdbPassword),
		option)
	defer client.Close()

	return
}
