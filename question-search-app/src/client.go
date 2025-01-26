package main

import (
    "context"
    "fmt"
    "log"
    "os"

    pb "github.com/bhargava-prashant/question-search-app/proto"
    "google.golang.org/grpc"
)

func main() {
    // Check if query is provided as command-line argument
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run client.go \"search query\"")
        os.Exit(1)
    }

    // Get search query from command-line argument
    query := os.Args[1]

    // Connect to the gRPC server
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Create a client for the QuestionService
    client := pb.NewQuestionServiceClient(conn)

    // Prepare the search request
    req := &pb.SearchRequest{Query: query}

    // Send the request and get the response
    resp, err := client.SearchQuestions(context.Background(), req)
    if err != nil {
        log.Fatalf("Error during SearchQuestions: %v", err)
    }

    // Print the results
    if len(resp.GetQuestions()) == 0 {
        fmt.Println("No questions found matching the query.")
        return
    }

    fmt.Printf("Search Results for query '%s':\n", query)
    for _, question := range resp.GetQuestions() {
        fmt.Printf("\n--- Question ---\n")
        fmt.Printf("ID: %s\n", question.GetId())
        fmt.Printf("Title: %s\n", question.GetTitle())
        fmt.Printf("Type: %s\n", question.GetType())

        // Handle type-specific details
        switch question.GetType() {
        case "MCQ":
            fmt.Println("Options:")
            for _, opt := range question.GetOptions() {
                fmt.Printf("- %s (Correct: %v)\n", opt.GetText(), opt.GetIsCorrectAnswer())
            }
        case "ANAGRAM":
            fmt.Println("Blocks:")
            for _, block := range question.GetBlocks() {
                fmt.Printf("- Text: %s (Show: %v, Answer: %v)\n", 
                    block.GetText(), block.GetShowInOption(), block.GetIsAnswer())
            }
            if question.GetSolution() != "" {
                fmt.Printf("Solution: %s\n", question.GetSolution())
            }
        }
    }
}