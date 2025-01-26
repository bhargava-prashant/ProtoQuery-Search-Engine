package main

import (
    "log"
    
    "github.com/bhargava-prashant/question-search-app/database"
    "github.com/bhargava-prashant/question-search-app/server"
)

func main() {
    connectionString := "mongodb+srv://prashantbhargava365:KSwWQMbpngg3kdtw@cluster0.2mhx1.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
    dbName := "questionsDB"
    
    mongoClient := database.NewMongoDBClient(connectionString, dbName)
    defer mongoClient.Disconnect()

    log.Println("Starting gRPC server on :50051")
    server.StartGRPCServer(mongoClient, ":50051")
}