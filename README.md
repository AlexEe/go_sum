# goSum : A service for mathematical operations on the CLI

## Overview 

goSum is a microservice written in Golang. A cli client communicates with the external server via gRPC. The client sends an array of ints which is entered via flag input on the command line. The server then calculates the sum of these numbers and returns the result which is printed out on the command line.

Check out the latest [Docker image](https://cloud.docker.com/repository/docker/alexeecode/gosum) here.

## Install

### Get source code from Github
```
go get github.com/AlexEe/goSum
```
### Run the server

#### With Docker image
```
docker run -p 8080:8080 alexeecode/gosum:1.0.1
```
#### Without Docker image
```
go run cmd/server/main.go
```
You can specify a port using the flag '-p' or 'port':
```
go run cmd/server/main.go -p 8080
```
The default port is 8080.
### Run the client
```
go run cmd/client/main.go
```
### Use the sum service
The sum service is started by using the 'sum' subcommand on the CLI,
followed by the flag '-n' or 'numbers' with an array of numbers, each
separated by ','.
```
./gosum_cli -n 1,2,3
```
If no numbers are provided, the CLI tool will prompt the user for input from the command line.
You can specify the url using the flag '-u' or 'url':
```
./gosum_cli -u localhost:8080
```
### Run the tests
```
go test ./...
```