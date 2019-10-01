# goSum : A service for mathematical operations on the CLI

## Overview [![Dockerhub]()](https://cloud.docker.com/repository/docker/alexeecode/gosum)

goSum is a microservice written in Golang. A cli client communicates with the external server via gRPC. The client sends an array of ints which is entered via flag input on the command line. The server then calculates the sum of these numbers and returns the result which is printed out on the command line.

## Install

### Get source code from Github
```
go get github.com/AlexEe/goSum
```
### Run the server

#### With Docker image
```
docker run alexee/gosum:1.0.0
```
You can specify a port using the flag '-p' or 'port':
```
docker run alexee/gosum:1.0.0 -p 8080
```
The default port is 8080.

#### Without Docker image
```
go run server/main.go
```
You can specify a port using the flag '-p' or 'port':
```
go run server/main.go -p 8080
```
The default port is 8080.
### Run the client
```
go run client/main.go
```
### Use the sum service
The sum service is started by using the 'sum' subcommand on the CLI,
followed by the flag '-n' or 'numbers' with an array of numbers, each
separated by ','.
```
go run client/main.go sum -n 1,2,3
```
You can specify the url using the flag '-u' or 'url':
```
go run client/main.go -u localhost:8080
```
### Run the tests
```
go test ./...
```