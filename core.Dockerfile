FROM  golang:1.22.4 as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o /handle ./cmd/core/main.go

# FROM scratch

# COPY --from=build /handle /handle
# COPY .env .env

EXPOSE 8080

ENTRYPOINT ["/handle"]