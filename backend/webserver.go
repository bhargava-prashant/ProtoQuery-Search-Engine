package main

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"encoding/json"
	"os"
	"github.com/bhargava-prashant/question-search-app/proto"
	"google.golang.org/grpc"
)

// searchQuestionsHandler handles the search query and interacts with the gRPC server
func searchQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // Assuming your server is running on localhost:50051
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect: %v", err), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Create a client for the QuestionService
	client := proto.NewQuestionServiceClient(conn)

	// Prepare the search request (adjust the query as needed)
	query := r.URL.Query().Get("query") // Get query parameter from the URL
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}

	req := &proto.SearchRequest{Query: query}

	// Send the request and get the response
	resp, err := client.SearchQuestions(context.Background(), req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error during SearchQuestions: %v", err), http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var result []map[string]string
	// Collect results
	for _, question := range resp.GetQuestions() {
		result = append(result, map[string]string{
			"ID":    question.GetId(),
			"Title": question.GetTitle(),
			"Type":  question.GetType(),
		})
	}

	// Encode to JSON
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
	}
}

// ServeIndexFile serves the index.html file directly when accessed from root URL
func ServeIndexFile(w http.ResponseWriter, r *http.Request) {
	// Open the index.html file from the "web" directory
	file, err := os.Open("./web/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error opening index.html: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Get file stats for the file's modification time
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting file stats: %v", err), http.StatusInternalServerError)
		return
	}

	// Serve the file content
	http.ServeContent(w, r, "index.html", fileInfo.ModTime(), file)
}

func main() {
	// Handle the root URL by serving the index.html file directly
	http.HandleFunc("/", ServeIndexFile)

	// Handle search requests
	http.HandleFunc("/search", searchQuestionsHandler)

	// Start the web server on port 8080
	log.Println("Web server starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}
