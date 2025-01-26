# Use an official Go image as the base
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all other files into the container
COPY . ./

# Build the server.go and webserver.go files
RUN go build -o server server.go
RUN go build -o webserver webserver.go

# Expose the ports used by your servers
EXPOSE 8080 8081

# Run both servers concurrently
CMD ["sh", "-c", "./server & ./webserver"]
