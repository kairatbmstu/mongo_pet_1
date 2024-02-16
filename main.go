package main

import (
	"context"
	"log"
	"mongo_peg_1/proto/grpcgen"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// // Set up MongoDB client
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	panic(err)
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Disconnect(ctx)

	// // Select database and collection
	// database := client.Database("your_database_name")
	// pageRepository := repository.NewPageRepository(database)

	// // Создание продюсера Kafka
	// producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	// if err != nil {
	// 	log.Fatalf("Failed to create producer: %v", err)
	// }
	// defer producer.Close()

	// var schedulerService = service.NewSchedulerService(pageRepository, producer)

	// var downloadService = service.DownloadService{
	// 	PageRepository: pageRepository,
	// }
	// var linkExtractService = service.LinkExtractService{
	// 	PageRepository: pageRepository,
	// }

	// downloadService.Listen()
	// linkExtractService.Listen()
	// schedulerService.Run()

	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register your service implementation
	var server = PageServiceServerImpl{}
	grpcgen.RegisterPageServiceServer(s, &server)

	// Register reflection service on gRPC server
	reflection.Register(s)

	log.Println("Server listening on port 50051")
	// Serve requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	time.Sleep(1000000 * time.Hour)
}

// UnimplementedPageServiceServer must be embedded to have forward compatible implementations.
type PageServiceServerImpl struct {
	grpcgen.UnimplementedPageServiceServer
}

func (PageServiceServerImpl) CreatePage(ctx context.Context, req *grpcgen.CreatePageRequest) (*grpcgen.CreatePageResponse, error) {
	response := grpcgen.CreatePageResponse{
		Id:       req.Id,
		Created:  req.Created,
		Updated:  req.Updated,
		Url:      req.Url,
		Status:   req.Status,
		Html:     req.Html,
		PageRank: 1.0,
		Links:    req.Links,
	}

	return &response, nil
}
func (PageServiceServerImpl) UpdatePage(ctx context.Context, req *grpcgen.UpdatePageRequest) (*grpcgen.UpdatePageResponse, error) {
	response := grpcgen.UpdatePageResponse{
		Id:       req.Id,
		Created:  req.Created,
		Updated:  req.Updated,
		Url:      req.Url,
		Status:   req.Status,
		Html:     req.Html,
		PageRank: 1.0,
		Links:    req.Links,
	}

	return &response, nil
}
func (PageServiceServerImpl) DeletePage(ctx context.Context, req *grpcgen.DeletePageRequest) (*grpcgen.DeletePageResponse, error) {
	response := grpcgen.DeletePageResponse{
		Id: req.Id,
	}

	return &response, nil
}
func (PageServiceServerImpl) FindById(ctx context.Context, req *grpcgen.FindByIdRequest) (*grpcgen.FindByIdResponse, error) {
	response := grpcgen.FindByIdResponse{
		Id:       "",
		Created:  nil,
		Updated:  nil,
		Url:      "",
		Status:   "",
		Html:     "",
		PageRank: 1.0,
	}

	return &response, nil
}
