FROM golang:1.16.3
ADD . /src/
WORKDIR /src
ENTRYPOINT GOOS=linux GOARCH=amd64 go test -tags=integration ./...