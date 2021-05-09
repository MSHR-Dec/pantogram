.PHONY: protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative datastore/pb/*.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative timeseries/pb/*.proto

.PHONY: build/api
build:
	docker compose build api
.PHONY: build/datastore
build:
	docker compose build datastore
.PHONY: build/timeseries
build:
	docker compose build timeseries

.PHONY: docker/up
docker/up:
	docker compose up
.PHONY: docker/down
docker-compose/down:
	docker compose down

.PHONY: database/up
database/up:
	docker compose up -d mysql influxdb
.PHONY: database/down
database/down:
	docker compose stop mysql influxdb

.PHONY: all
all: protoc build/api build/datastore build/timeseries docker/up
