docker build . -t rent:local

docker-compose -f docker-compose.yaml up -d

go test -tags=integration ./...

docker-compose -f docker-compose.yaml -f dev/docker-compose.local.yaml down