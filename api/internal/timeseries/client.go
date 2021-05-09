package timeseries

import (
	"context"
	"github.com/MSHR-Dec/pantogram/timeseries/pb"
	"google.golang.org/grpc"
	"log"
)

type TimeseriesApi struct {
	conn *grpc.ClientConn
}

func NewTimeseriesApi(c *grpc.ClientConn) *TimeseriesApi {
	return &TimeseriesApi{
		conn: c,
	}
}

func (t TimeseriesApi) Request(req *pb.SaveRequest) {
	c := pb.NewTimeseriesClient(t.conn)
	_, err := c.SaveTimeseries(context.TODO(), req)
	if err != nil {
		log.Println(err)
	}
}
