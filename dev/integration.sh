docker-compose -f docker-compose.yaml up -d --force-recreate -V

# database container isn't always ready after spin up and can cause EOF errors...
sleep 1

go test -tags=integration ./...

docker-compose -f docker-compose.yaml down