package main

import (
	"log"
	"net"

	health "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/MSHR-Dec/pantogram/datastore/application"
	"github.com/MSHR-Dec/pantogram/datastore/infrastructure"
	"github.com/MSHR-Dec/pantogram/datastore/infrastructure/rdb"
	"github.com/MSHR-Dec/pantogram/datastore/pb"
	"github.com/MSHR-Dec/pantogram/datastore/server"
	"google.golang.org/grpc"
)

func main() {
	db := infrastructure.InitMysql()

	routeRepository := rdb.NewRouteRepository(db)
	companyRepository := rdb.NewCompanyRepository(db)
	prefectureRepository := rdb.NewPrefectureRepository(db)

	routeDetailInteractor := application.NewRouteDetailInteractor(
		routeRepository,
		companyRepository,
		prefectureRepository)

	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterDatastoreServer(s, &server.Datastore{
		RouteDetail: routeDetailInteractor,
	})

	health.RegisterHealthServer(s, &server.HealthHandler{})

	log.Println("Start gRPC ...")
	s.Serve(listenPort)
}
