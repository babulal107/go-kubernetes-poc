# Go Kubernetes POC
Go Lang Application and deploy it on Kubernetes 

## Makefile
- `make build` to build the golang binary of application
- `make run` to run application on local
- `make clean` to clean generated binary of application

## Run on local
- `make run`

## Unit Testing
- `go test -short ./...` to run without integration tests
- `go test -v ./...` to run with integration tests

## Health Check API
- `http://localhost:8080/health`

## Build docker image
- `docker build -t babulal107/go-kubernetes-poc:latest .`

## Run docker Container 
- `docker run -p 8090:8080 -d babulal107/go-kubernetes-poc`

## Push docker to docker hub repository

- `docker push babulal107/go-kubernetes-poc`