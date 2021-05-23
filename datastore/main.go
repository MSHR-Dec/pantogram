package main

import (
	"log"
	"net"

	"github.com/opentracing/opentracing-go"

	"github.com/MSHR-Dec/pantogram/api/pkg/jaeger"

	health "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/MSHR-Dec/pantogram/datastore/application"
	"github.com/MSHR-Dec/pantogram/datastore/infrastructure"
	"github.com/MSHR-Dec/pantogram/datastore/infrastructure/rdb"
	"github.com/MSHR-Dec/pantogram/datastore/pb"
	"github.com/MSHR-Dec/pantogram/datastore/server"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
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

	tracer, closer, err := jaeger.NewTracer()
	defer closer.Close()
	if err != nil {
		log.Fatalln(err)
	}
	opentracing.SetGlobalTracer(tracer)

	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			// add opentracing stream interceptor to chain
			grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			// add opentracing unary interceptor to chain
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		)),
	)
	pb.RegisterDatastoreServer(s, &server.Datastore{
		RouteDetail: routeDetailInteractor,
	})

	health.RegisterHealthServer(s, &server.HealthHandler{})

	log.Println("Start gRPC ...")
	s.Serve(listenPort)
}
