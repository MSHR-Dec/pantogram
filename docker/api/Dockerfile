FROM golang:1.15.8-alpine3.13 as build

WORKDIR /api
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./api/cmd

FROM alpine

COPY --from=build /api/main /main

CMD ["/main"]