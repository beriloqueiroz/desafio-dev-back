FROM  golang:1.22.4 as build
ENV TZ="America/Fortaleza"
WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o /handle ./cmd/main.go

# FROM scratch

# COPY --from=build /handle /handle
# COPY .env .env

ENTRYPOINT ["/handle"]