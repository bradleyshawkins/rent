docker build . -t rent:local

docker-compose -f docker-compose.yaml -f dev/docker-compose.local.yaml up -d --force-recreate -V

go test -tags=integration ./...

docker-compose -f docker-compose.yaml -f dev/docker-compose.local.yaml down