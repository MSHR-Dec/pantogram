.PHONY: protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative datastore/pb/*.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative timeseries/pb/*.proto

.PHONY: build/api
build/api:
	docker compose build api
.PHONY: build/datastore
build/datastore:
	docker compose build datastore
.PHONY: build/timeseries
build/timeseries:
	docker compose build timeseries
.PHONY: build/mysql
build/mysql:
	docker compose build mysql

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

.PHONY: push-private/api
push-private/api:
	docker tag pantogram_api 192.168.0.10:5000/api:v0.0.0
	docker push 192.168.0.10:5000/api:v0.0.0
.PHONY: push-private/datastore
push-private/datastore:
	docker tag pantogram_datastore 192.168.0.10:5000/datastore:v0.0.2
	docker push 192.168.0.10:5000/datastore:v0.0.2
.PHONY: push-private/timeseries
push-private/timeseries:
	docker tag pantogram_timeseries 192.168.0.10:5000/timeseries:v0.0.3
	docker push 192.168.0.10:5000/timeseries:v0.0.3
.PHONY: push-private/mysql
push-private/mysql:
	docker tag pantogram_mysql 192.168.0.10:5000/mysql:v0.0.1
	docker push 192.168.0.10:5000/mysql:v0.0.1

.PHONY: push-private
push-private: push-private/api push-private/datastore push-private/timeseries push-private/mysql
