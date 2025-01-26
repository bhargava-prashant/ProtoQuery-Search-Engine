package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/bhargava-prashant/question-search-app/database"
    "github.com/bhargava-prashant/question-search-app/proto"
    "github.com/bhargava-prashant/question-search-app/server"
    "github.com/gorilla/handlers"
    "google.golang.org/grpc"
)

// CORS Middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight requests
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next(w, r)
    }
}

func searchQuestionsHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Received search request: %s", r.URL.String())

    // Connect to the gRPC server
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Printf("Failed to connect to gRPC server: %v", err)
        http.Error(w, fmt.Sprintf("Failed to connect: %v", err), http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    // Create a client for the QuestionService
    client := proto.NewQuestionServiceClient(conn)

    // Prepare the search request
    query := r.URL.Query().Get("query")
    if query == "" {
        log.Printf("Missing query parameter")
        http.Error(w, "Missing query parameter", http.StatusBadRequest)
        return
    }

    req := &proto.SearchRequest{Query: query}

    // Send the request and get the response
    resp, err := client.SearchQuestions(context.Background(), req)
    if err != nil {
        log.Printf("Error during SearchQuestions: %v", err)
        http.Error(w, fmt.Sprintf("Error during SearchQuestions: %v", err), http.StatusInternalServerError)
        return
    }

    // Write JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)

    // Prepare a more detailed result structure
    var result []map[string]interface{}
    for _, question := range resp.GetQuestions() {
        questionDetail := map[string]interface{}{
            "ID":    question.GetId(),
            "Title": question.GetTitle(),
            "Type":  question.GetType(),
        }

        switch question.GetType() {
        case "MCQ":
            var options []map[string]interface{}
            for _, opt := range question.GetOptions() {
                options = append(options, map[string]interface{}{
                    "text":            opt.GetText(),
                    "isCorrectAnswer": opt.GetIsCorrectAnswer(),
                })
            }
            questionDetail["options"] = options
        
        case "ANAGRAM":
            var blocks []map[string]interface{}
            for _, block := range question.GetBlocks() {
                blocks = append(blocks, map[string]interface{}{
                    "text":         block.GetText(),
                    "showInOption": block.GetShowInOption(),
                    "isAnswer":     block.GetIsAnswer(),
                })
            }
            questionDetail["blocks"] = blocks
            questionDetail["solution"] = question.GetSolution()
        }

        result = append(result, questionDetail)
    }

    if err := json.NewEncoder(w).Encode(result); err != nil {
        log.Printf("Error encoding JSON: %v", err)
        http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
    }

    log.Printf("Successfully returned %d results", len(result))
}

func ServeIndexFile(w http.ResponseWriter, r *http.Request) {
    file, err := os.Open("./web/index.html")
    if err != nil {
        log.Printf("Error opening index.html: %v", err)
        http.Error(w, fmt.Sprintf("Error opening index.html: %v", err), http.StatusInternalServerError)
        return
    }
    defer file.Close()

    fileInfo, err := file.Stat()
    if err != nil {
        log.Printf("Error getting file stats: %v", err)
        http.Error(w, fmt.Sprintf("Error getting file stats: %v", err), http.StatusInternalServerError)
        return
    }

    http.ServeContent(w, r, "index.html", fileInfo.ModTime(), file)
}

func main() {
    // MongoDB connection
    connectionString := "mongodb+srv://prashantbhargava365:KSwWQMbpngg3kdtw@cluster0.2mhx1.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
    dbName := "questionsDB"
    
    mongoClient := database.NewMongoDBClient(connectionString, dbName)
    defer mongoClient.Disconnect()

    // Start gRPC server in a goroutine
    go func() {
        log.Println("Starting gRPC server on :50051")
        server.StartGRPCServer(mongoClient, ":50051")
    }()

    // Set up HTTP handlers with CORS middleware
    http.HandleFunc("/", enableCORS(ServeIndexFile))
    http.HandleFunc("/search", enableCORS(searchQuestionsHandler))

    // Get port from environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Create custom CORS handler
    corsMiddleware := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )

    // Start the server with both custom middleware and gorilla/handlers CORS
    log.Printf("Web server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(http.DefaultServeMux)))
}