FROM golang:1.16.3 AS build
ADD . /src/
WORKDIR /src/cmd/rent
RUN GOOS=linux GOARCH=amd64 go build -o rent



FROM alpine
RUN apk add --no-cache \
        perl \
        wget \
        openssl \
        ca-certificates \
        libc6-compat \
        libstdc++

WORKDIR /app
ADD mysql/schema /app/schema
COPY --from=build /src/cmd/rent/rent /app/
ENTRYPOINT ./rent