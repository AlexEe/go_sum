# Import golang base image
FROM golang:1.13 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sum_server server/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/sum_server .

CMD [ "./sum_server" ]

# Expose port 8080 to the outside world
EXPOSE 8080