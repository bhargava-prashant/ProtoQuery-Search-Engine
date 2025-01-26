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
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run client.go \"search query\"")
        os.Exit(1)
    }
    query := os.Args[1]
    // gRPC server
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()
    client := pb.NewQuestionServiceClient(conn)
    req := &pb.SearchRequest{Query: query}

    resp, err := client.SearchQuestions(context.Background(), req)
    if err != nil {
        log.Fatalf("Error during SearchQuestions: %v", err)
    }

    // results
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
