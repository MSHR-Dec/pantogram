package tsdb

import (
	"context"
	"fmt"
	"time"

	"github.com/MSHR-Dec/pantogram/timeseries/domain/repository"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type CompanyRepository struct {
	db       influxdb2.Client
	database string
}

func NewCompanyRepository(d influxdb2.Client, name string) repository.CompanyRepository {
	return &CompanyRepository{
		db:       d,
		database: name,
	}
}

func (r *CompanyRepository) Save(t map[string]string, f map[string]interface{}, n time.Time) {
	writeAPI := r.db.WriteAPIBlocking("", fmt.Sprintf("%s/autogen", r.database))
	p := influxdb2.NewPoint("company", t, f, n)
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Printf("Write error: %s\n", err.Error())
	}
}
