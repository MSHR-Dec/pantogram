package server

import (
	"context"
	"github.com/MSHR-Dec/pantogram/timeseries/application"
	"github.com/MSHR-Dec/pantogram/timeseries/pb"
)

type Timeseries struct {
	pb.UnimplementedTimeseriesServer
	TimeseriesApplication application.Timeseries
}

func (t Timeseries) SaveTimeseries(ctx context.Context, req *pb.SaveRequest) (*pb.SaveResponse, error) {
	var routes []application.Route
	var companies []application.Company
	var prefectures []application.Prefecture

	for _, detail := range req.RouteDetails {
		prefNames := func(prefs []*pb.Prefecture) (names []string) {
			for _, prefecture := range prefs {
				names = append(names, prefecture.Name)
			}
			return
		}

		routes = append(routes, application.Route{
			Name:     detail.Name,
			ComName:  detail.Company,
			PrefName: prefNames(detail.Prefectures),
		})

		companies = append(companies, application.Company{Name: detail.Company})

		for _, prefecture := range detail.Prefectures {
			prefectures = append(prefectures, application.Prefecture{
				Name:  prefecture.Name,
				Count: int(prefecture.Count),
			})
		}
	}

	t.TimeseriesApplication.SaveTimeseries(routes, companies, prefectures)

	return &pb.SaveResponse{}, nil
}
