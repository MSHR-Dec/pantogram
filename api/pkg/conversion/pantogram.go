package conversion

import (
	dspb "github.com/MSHR-Dec/pantogram/datastore/pb"
	tspb "github.com/MSHR-Dec/pantogram/timeseries/pb"
)

func ConvertDsToTs(dsStruct []*dspb.RouteDetail) *tspb.SaveRequest {
	var tsRouteDetails []*tspb.RouteDetail

	for _, ds := range dsStruct {
		var tsPrefectures []*tspb.Prefecture

		for _, dsPrefecture := range ds.Prefectures {
			tsPrefecture := &tspb.Prefecture{
				Name:  dsPrefecture.Name,
				Count: dsPrefecture.Count,
			}
			tsPrefectures = append(tsPrefectures, tsPrefecture)
		}

		tsRouteDetail := &tspb.RouteDetail{
			Name:        ds.Name,
			Company:     ds.Company,
			Prefectures: tsPrefectures,
		}
		tsRouteDetails = append(tsRouteDetails, tsRouteDetail)
	}

	return &tspb.SaveRequest{RouteDetails: tsRouteDetails}
}
