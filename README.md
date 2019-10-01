# goSum : A service for mathematical operations on the CLI

## Overview [![GoDoc](https://godoc.org/github.com/AlexEe/goSum?status.svg)](https://godoc.org/github.com/AlexEe/goSum)

goSum is a microservice written in Golang. A cli client communicates with the external server via gRPC. The client sends an array of ints which is entered via flag input on the command line. The server then calculates the sum of these numbers and returns the result which is printed out on the command line.

## Install

### Get source code from Github
```
go get github.com/AlexEe/goSum
```
### Run the server

#### With Docker
```
docker run alexee/gosum:1.0.0
```
It is possible to specify a port using the flag '-p' or 'port':
```
docker run alexee/gosum:1.0.0 -p 8080
```
The default port is 8080.

#### Without Docker
```
go run server/main.go
```
It is possible to specify a port using the flag '-p' or 'port':
```
go run server/main.go -p 8080
```
The default port is 8080.
### Run the client
```
go run client/main.go
```
It is possible to specify the url using the flag '-u' or 'url':
```
go run client/main.go -u localhost:8080
```
### Use the sum service
The sum service is started by using the 'sum' subcommand on the CLI,
followed by the flag '-n' or 'numbers' with an array of numbers, each
separated by ','.
```
go run client/main.go sum -n 1,2,3
```