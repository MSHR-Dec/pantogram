package main

import (
	"log"
	"net"

	"github.com/MSHR-Dec/pantogram/timeseries/application"
	"github.com/MSHR-Dec/pantogram/timeseries/infrastructure"
	"github.com/MSHR-Dec/pantogram/timeseries/infrastructure/tsdb"
	"github.com/MSHR-Dec/pantogram/timeseries/pb"
	"github.com/MSHR-Dec/pantogram/timeseries/server"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	db := infrastructure.InitInfluxdb()

	routeRepository := tsdb.NewRouteRepository(db, infrastructure.Environment.InfluxdbDbname)
	companyRepository := tsdb.NewCompanyRepository(db, infrastructure.Environment.InfluxdbDbname)
	prefectureRepository := tsdb.NewPrefectureRepository(db, infrastructure.Environment.InfluxdbDbname)

	timeseriesApplication := application.NewTimeseriesInteractor(
		routeRepository,
		companyRepository,
		prefectureRepository)

	listenPort, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterTimeseriesServer(s, &server.Timeseries{
		TimeseriesApplication: timeseriesApplication,
	})

	health.RegisterHealthServer(s, &server.HealthHandler{})

	log.Println("Start gRPC ...")
	s.Serve(listenPort)
}
