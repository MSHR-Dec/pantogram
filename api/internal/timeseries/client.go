package timeseries

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/MSHR-Dec/pantogram/timeseries/pb"
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
