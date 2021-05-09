package main

import (
	"log"
	"runtime"

	"github.com/carlescere/scheduler"

	"github.com/MSHR-Dec/pantogram/api/internal/timeseries"
	"github.com/MSHR-Dec/pantogram/api/pkg/conversion"

	"github.com/MSHR-Dec/pantogram/api/internal/datastore"
	"github.com/MSHR-Dec/pantogram/api/pkg/delay"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

type osEnvironment struct {
	Datastore  string `default:"localhost:8080"`
	Timeseries string `default:"localhost:8081"`
}

var environment osEnvironment

func init() {
	envconfig.Process("", &environment)
}

func main() {
	job := func() {
		var api delay.Delay
		routes := api.GetDelayRoutes()

		dsConn := connectDatastore()
		defer dsConn.Close()

		datastoreClient := datastore.NewDatastoreApi(dsConn)
		result := datastoreClient.Request(routes)

		tsdbConn := connectTimeseries()
		defer tsdbConn.Close()

		timeseriesClient := timeseries.NewTimeseriesApi(tsdbConn)
		timeseriesClient.Request(conversion.ConvertDsToTs(result))
	}
	scheduler.Every(1).Minutes().Run(job)
	runtime.Goexit()
}

func connectDatastore() (conn *grpc.ClientConn) {
	address := environment.Datastore

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("client connection error:", err)
	}

	return
}

func connectTimeseries() (conn *grpc.ClientConn) {
	address := environment.Timeseries

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("client connection error:", err)
	}

	return
}
