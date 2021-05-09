package server

import (
	"context"
	"errors"

	"github.com/MSHR-Dec/pantogram/datastore/application"
	"github.com/MSHR-Dec/pantogram/datastore/pb"
)

type Datastore struct {
	pb.UnimplementedDatastoreServer
	RouteDetail application.RouteDetail
}

func (d Datastore) GetRouteDetail(ctx context.Context, route *pb.Route) (*pb.RouteDetail, error) {
	if route.Name == "" {
		return nil, errors.New("the name of route not found")
	}

	var prefectures []*pb.Prefecture

	output, err := d.RouteDetail.GetRouteDetail(route.Name)
	if err != nil {
		return nil, err
	}

	for _, v := range output.Prefectures {
		prefecture := pb.Prefecture{
			Name:  v.Name,
			Count: v.RouteCount,
		}
		prefectures = append(prefectures, &prefecture)
	}

	return &pb.RouteDetail{
		Name:        route.Name,
		Company:     output.Company.Name,
		Prefectures: prefectures,
	}, nil
}
