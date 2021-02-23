package persistence

import (
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InitInfluxdb(env string, user string, password string, host string) (client influxdb2.Client) {
	option := influxdb2.DefaultOptions()
	if env != "prd" {
		option.SetLogLevel(3)
	}

	client = influxdb2.NewClientWithOptions(
		fmt.Sprintf("http://%s", host),
		fmt.Sprintf("%s:%s", user, password),
		option)
	defer client.Close()

	return
}
