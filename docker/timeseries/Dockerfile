FROM golang:1.15.8-alpine3.13 as build

WORKDIR /timeseries
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./timeseries

FROM alpine

COPY --from=build /timeseries/main /main
RUN apk update \
    && apk add curl

CMD ["/main"]