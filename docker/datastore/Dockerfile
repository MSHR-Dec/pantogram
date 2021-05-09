FROM golang:1.15.8-alpine3.13 as build

WORKDIR /datastore
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./datastore

FROM alpine

COPY --from=build /datastore/main /main
RUN apk update \
    && apk add mysql-client

COPY ./docker/datastore/script/start.sh /start.sh
RUN chmod 755 /start.sh

CMD ["/start.sh"]