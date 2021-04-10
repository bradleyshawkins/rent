FROM golang:1.16.3 AS build
ENV GOOS=linux
ADD . /src
RUN cd /src && go build -o rent



FROM alpine
WORKDIR /app
ADD mysql/schema /app
COPY --from=build /src/rent /app/
ENTRYPOINT ./rent