package server

import (
    "context"
    "log"
    "net"

    "github.com/bhargava-prashant/question-search-app/database"
    pb "github.com/bhargava-prashant/question-search-app/proto"

    "go.mongodb.org/mongo-driver/bson"
    "google.golang.org/grpc"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionServer struct {
    pb.UnimplementedQuestionServiceServer
    mongoClient *database.MongoDBClient
}

func NewQuestionServer(mongoClient *database.MongoDBClient) *QuestionServer {
    return &QuestionServer{
        mongoClient: mongoClient,
    }
}

func (s *QuestionServer) SearchQuestions(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
    log.Printf("Received Search Query: %v", req.Query)  // Log the search query received

    collection := s.mongoClient.DB.Collection("questions")
    
    // Create a filter to search by title or type using regex
    filter := bson.M{
        "$or": []bson.M{
            {"title": bson.M{"$regex": req.Query, "$options": "i"}},  // Case-insensitive match on title
            {"type": bson.M{"$regex": req.Query, "$options": "i"}},   // Case-insensitive match on type
        },
    }

    // Perform the database query
    cur, err := collection.Find(ctx, filter)
    if err != nil {
        log.Printf("Error during query execution: %v", err)  // Log any errors
        return nil, err
    }
    defer cur.Close(ctx)

    var questions []*pb.Question
    for cur.Next(ctx) {
        var question map[string]interface{}
        if err := cur.Decode(&question); err != nil {
            log.Printf("Error decoding question: %v", err)
            continue
        }
    
        // Convert the ObjectID to string using Hex() method
        pbQuestion := &pb.Question{
            Id:    question["_id"].(primitive.ObjectID).Hex(), // ObjectID to string
            Type:  question["type"].(string),
            Title: question["title"].(string),
        }
    
        log.Printf("Found Question: %v", pbQuestion)
    
        questions = append(questions, pbQuestion)
    }
    

    log.Printf("Returning %d questions", len(questions))  // Log the number of questions found
    return &pb.SearchResponse{Questions: questions}, nil
}


func StartGRPCServer(mongoClient *database.MongoDBClient, port string) {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    
    grpcServer := grpc.NewServer()
    questionServer := NewQuestionServer(mongoClient)
    pb.RegisterQuestionServiceServer(grpcServer, questionServer)
    
    log.Printf("gRPC server listening on %v", port)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}