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
    "go.mongodb.org/mongo-driver/mongo/options"
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
    log.Printf("Received Search Query: %v", req.Query)

    collection := s.mongoClient.DB.Collection("questions")
    
    // Create a filter to search by title or type using regex
    filter := bson.M{
        "$or": []bson.M{
            {"title": bson.M{"$regex": req.Query, "$options": "i"}},
            {"type": bson.M{"$regex": req.Query, "$options": "i"}},
        },
    }

    // Set up pagination (limit to 50 results)
    findOptions := options.Find().SetLimit(50)

    // Perform the database query
    cur, err := collection.Find(ctx, filter, findOptions)
    if err != nil {
        log.Printf("Error during query execution: %v", err)
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
    
        // Convert the ObjectID to string
        pbQuestion := &pb.Question{
            Id:    question["_id"].(primitive.ObjectID).Hex(),
            Type:  question["type"].(string),
            Title: question["title"].(string),
        }
    
        // Handle different question types
        switch pbQuestion.Type {
        case "MCQ":
            if options, ok := question["options"].(primitive.A); ok {
                for _, opt := range options {
                    optMap := opt.(map[string]interface{})
                    pbQuestion.Options = append(pbQuestion.Options, &pb.Option{
                        Text:            optMap["text"].(string),
                        IsCorrectAnswer: optMap["isCorrectAnswer"].(bool),
                    })
                }
            }
        case "ANAGRAM":
            if blocks, ok := question["blocks"].(primitive.A); ok {
                for _, block := range blocks {
                    blockMap := block.(map[string]interface{})
                    pbQuestion.Blocks = append(pbQuestion.Blocks, &pb.Block{
                        Text:         blockMap["text"].(string),
                        ShowInOption: blockMap["showInOption"].(bool),
                        IsAnswer:     blockMap["isAnswer"].(bool),
                    })
                }
                if solution, ok := question["solution"].(string); ok {
                    pbQuestion.Solution = solution
                }
            }
        }
    
        questions = append(questions, pbQuestion)
    }
    
    log.Printf("Returning %d questions", len(questions))
    return &pb.SearchResponse{Questions: questions}, nil
}

func StartGRPCServer(mongoClient *database.MongoDBClient, port string) {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    
    // Increase max message size
    grpcServer := grpc.NewServer(
        grpc.MaxRecvMsgSize(50*1024*1024),  // 50 MB
        grpc.MaxSendMsgSize(50*1024*1024),  // 50 MB
    )
    
    questionServer := NewQuestionServer(mongoClient)
    pb.RegisterQuestionServiceServer(grpcServer, questionServer)
    
    log.Printf("gRPC server listening on %v", port)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}