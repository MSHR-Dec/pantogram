package datastore

import (
	"context"
	"log"

	"github.com/MSHR-Dec/pantogram/datastore/pb"
	"google.golang.org/grpc"
)

type DatastoreApi struct {
	conn *grpc.ClientConn
}

func NewDatastoreApi(c *grpc.ClientConn) *DatastoreApi {
	return &DatastoreApi{
		conn: c,
	}
}

func (d DatastoreApi) Request(routes []string) []*pb.RouteDetail {
	var output []*pb.RouteDetail

	datastoreCh := make(chan *pb.RouteDetail, len(routes))
	defer close(datastoreCh)

	for _, name := range routes {
		go d.requestDatastoreTask(datastoreCh, d.conn, name)
		result := <-datastoreCh
		if result != nil {
			output = append(output, result)
		}
	}

	return output
}

func (d DatastoreApi) requestDatastoreTask(datastoreCh chan<- *pb.RouteDetail, conn *grpc.ClientConn, routeName string) {
	c := pb.NewDatastoreClient(conn)

	request := &pb.Route{
		Name: routeName,
	}

	res, err := c.GetRouteDetail(context.TODO(), request)
	if err != nil {
		log.Println(err)
	}

	datastoreCh <- res
}
