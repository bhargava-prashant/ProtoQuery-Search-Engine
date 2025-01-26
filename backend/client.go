package main

import (
    "context"
    "fmt"
    "log"

    pb "github.com/bhargava-prashant/question-search-app/proto"
    "google.golang.org/grpc"
)

func main() {
    // Connect to the gRPC server
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // Assuming your server is running on localhost:50051
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Create a client for the QuestionService
    client := pb.NewQuestionServiceClient(conn)

    // Prepare the search request (adjust the query as needed)
    req := &pb.SearchRequest{Query: "Which word is a synonym for 'Limit' in the context"}

    // Send the request and get the response
    resp, err := client.SearchQuestions(context.Background(), req)
    if err != nil {
        log.Fatalf("Error during SearchQuestions: %v", err)
    }

    // Print the results
    fmt.Println("Search Results:")
    for _, question := range resp.GetQuestions() {
        fmt.Printf("ID: %s, Title: %s, Type: %s\n", question.GetId(), question.GetTitle(), question.GetType())
    }
}
