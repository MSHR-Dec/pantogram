FROM golang:1.15.8-alpine3.13 as build

WORKDIR /timeseries
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./timeseries

FROM alpine

COPY --from=build /timeseries/main /main
RUN apk update && \
    apk add curl && \
    apk add wget
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

CMD ["/main"]
