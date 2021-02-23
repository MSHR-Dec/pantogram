FROM golang:1.15.8-alpine3.13 as build

WORKDIR /pantogram
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./src

FROM alpine

COPY --from=build /pantogram/main /main
RUN apk update \
    && apk add mysql-client

CMD ["/main"]