
FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o server server.go
RUN go build -o webserver webserver.go

EXPOSE 8080 8081

CMD ["sh", "-c", "./server & ./webserver"]
