package main

import (
	"runtime"

	"github.com/MSHR-Dec/pantogram/src/config"
	"github.com/MSHR-Dec/pantogram/src/infrastructure/persistence"
	"github.com/MSHR-Dec/pantogram/src/infrastructure/persistence/rdb"
	"github.com/MSHR-Dec/pantogram/src/infrastructure/persistence/tsdb"
	"github.com/MSHR-Dec/pantogram/src/interfaces/client"
	"github.com/MSHR-Dec/pantogram/src/usecase"
	"github.com/carlescere/scheduler"
)

func main() {
	job := func() {
		delays := new(client.Delays)
		delays.Run()

		pantogram := inject()
		pantogram.Create(delays.Names)
	}
	scheduler.Every(1).Minutes().Run(job)
	runtime.Goexit()
}

func inject() usecase.PantogramUsecase {
	mysqlConn := persistence.InitMysql(
		config.Environment.Env,
		config.Environment.MysqlUser,
		config.Environment.MysqlPassword,
		config.Environment.MysqlUrl,
		config.Environment.MysqlDbname)
	influxdbConn := persistence.InitInfluxdb(
		config.Environment.Env,
		config.Environment.InfluxdbUser,
		config.Environment.InfluxdbPassword,
		config.Environment.InfluxdbUrl)

	rdbRouteRepository := rdb.NewRouteRDBRepository(mysqlConn)
	rdbCompanyRepository := rdb.NewCompanyRDBRepository(mysqlConn)
	rdbPrefectureRepository := rdb.NewPrefectureRDBRepository(mysqlConn)

	tsdbRouteRepository := tsdb.NewRouteTSDBRepository(influxdbConn, config.Environment.InfluxdbDbname)
	tsdbCompanyRepository := tsdb.NewCompanyTSDBRepository(influxdbConn, config.Environment.InfluxdbDbname)
	tsdbPrefectureRepository := tsdb.NewPrefectureTSDBRepository(influxdbConn, config.Environment.InfluxdbDbname)

	return usecase.NewPantogramInteractor(
		rdbRouteRepository,
		rdbCompanyRepository,
		rdbPrefectureRepository,
		tsdbRouteRepository,
		tsdbCompanyRepository,
		tsdbPrefectureRepository)
}
