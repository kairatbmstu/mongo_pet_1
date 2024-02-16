package main

import (
	"context"
	"log"
	"mongo_peg_1/repository"
	"mongo_peg_1/service"
	"time"

	"github.com/IBM/sarama"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Set up MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	// Select database and collection
	database := client.Database("your_database_name")
	pageRepository := repository.NewPageRepository(database)

	// Создание продюсера Kafka
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	var schedulerService = service.NewSchedulerService(pageRepository, producer)

	var downloadService = service.DownloadService{
		PageRepository: pageRepository,
	}
	var linkExtractService = service.LinkExtractService{
		PageRepository: pageRepository,
	}

	downloadService.Listen()
	linkExtractService.Listen()
	schedulerService.Run()

	time.Sleep(1000000 * time.Hour)
}
