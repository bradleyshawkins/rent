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
EXPOSE 8080
ADD mysql/schema /app/schema
COPY --from=build /src/cmd/rent/rent /app/
HEALTHCHECK CMD wget http://localhost:8080/health || exit 1
ENTRYPOINT ./rent