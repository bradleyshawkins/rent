
# Rent

Rent is used to help Property owners find great tenants and simplifies communication between tenant and landlord.

## Installing / Getting started

Running the following commands will download, and begin the service
```shell
go get github.com/bradleyshawkins/rent
make deps.up
go run services/rent/main.go
```
* `go get github.com/bradleyshawkins/rent` Downloads the repository to `$GOPATH/src/github.com/bradleyshawkins/rent`
* `make deps.up` spins up a postgres image in a docker container
* `go run cmd/rent/main.go` starts the service

## Developing
Rent uses `make`, `go` and `docker-compose`. All 3 will need to be installed before beginning development


### Project structure
The following describes the purpose of each directory
* `rent` contains the service objects and methods
* `rent/cmd` contains the executables
* `rent/config` contains service configuration
* `rent/dev` contains developer scripts
* `rent/postgres` contains the postgres database logic
* `rent/rest` contains the rest endpoints

### Makefile
The Makefile contains operations used to help run the service
* `make deps.up` spins up all external resources (postgres)
* `make deps.down` spins down all external resources
* `make test.unit` runs unit tests
* `make test.integration` runs integration tests, but does not spin up external resources
* `make test.integration.env` spins up all external resources, builds the service in a container and runs the integration tests
* `service.container.build` builds the service docker image and tags it as local
* `env.start` starts all external services and the rent service in docker containers using docker compose
* `env.stop` stops all external services and the rent service and cleans up the volumes

### Building
The project can be built by running
```shell
cd services/rent/
go build
```

### Testing
Both unit tests and integration tests are used in Rent. Both are used using the Go testing framework. 
#### Unit Tests
Unit tests are written around business logic and don't require mocks or any external resources to be spun up.
#### Integration Tests
Integration tests are written around the rest endpoints in the `rest` package. They don't have any access into the internal packages of the service. Any operation they want to perform has to be done through endpoint requests.

Integration tests have the following line at the top of the file.
```
// +build integration

```
This makes it so integrations tests are only ran with the following command
`go test -tags=integration ./...`. The command exists in the Makefile with the following recipe. `make test.integration`

Integration tests will spin up all external resources in docker containers. They are all added to the same network which allows them to all communicate together by referencing the docker container name.

### Deploying / Publishing

Rent uses CI/CD via Github Actions. Both unit tests and integration tests are ran when code is pushed to Github. If both tests pass, the service is pushed to Heroku and is immediately live.

## Features

What's all the bells and whistles this project can perform?
* Users
  * Register
  * Cancel
  * Load
* Properties
  * Register
  * Remove
  * Load
  
## Contributing

I'm primarily writing this as an experiment on how to write and run services completely solo. If you'd like to contribute, please fork the repository and use a feature branch.
