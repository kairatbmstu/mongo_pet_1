package main

import (
	"log"
	"mongo_peg_1/repository"
	"mongo_peg_1/service"
	"time"

	"github.com/IBM/sarama"
)

func main() {

	// Создание продюсера Kafka
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	var pageRepository = repository.PageRepository{}
	var downloadService = service.DownloadService{
		PageRepository: pageRepository,
	}
	var linkExtractService = service.LinkExtractService{
		PageRepository: pageRepository,
	}
	var schedulerService = service.SchedulerService{
		PageRepository: pageRepository,
		SyncProducer:   producer,
	}

	downloadService.Listen()
	linkExtractService.Listen()
	schedulerService.Run()

	time.Sleep(1000000 * time.Hour)
}
