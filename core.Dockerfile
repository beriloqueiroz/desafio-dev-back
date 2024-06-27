FROM  golang:1.22.4 as build

USER root

WORKDIR /app

ADD cache_sync /cache_sync
ADD core /core
COPY ./core/.env .env
COPY ./go.work.core /go.work

RUN go work sync

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o /handle /core/cmd/main.go

# FROM scratch

# COPY --from=build /handle /handle
# COPY .env .env

EXPOSE 8080

ENTRYPOINT ["/handle"]