# 1. Use an official Golang base image
FROM golang:1.23.3

# 2. Set the working directory in the container
WORKDIR /app

# 3. Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# 4. Download Go module dependencies
RUN go mod download

# 5. Copy the application code
COPY . .

# 6. Build the Go application
RUN go build -o main .

# 7. Command to run the application
CMD ["./main"]