package service

import (
	"encoding/json"
	"log"
	"mongo_peg_1/model"
	"mongo_peg_1/repository"

	"github.com/IBM/sarama"
)

type SchedulerService struct {
	SyncProducer   sarama.SyncProducer
	PageRepository *repository.PageRepository
	Pages          []model.Page
}

func NewSchedulerService(repo *repository.PageRepository, SyncProducer sarama.SyncProducer) *SchedulerService {
	var service = SchedulerService{
		PageRepository: repo,
		SyncProducer:   SyncProducer,
	}
	service.Pages = []model.Page{
		{Url: "https://tengrinews.kz", Status: model.PageInitial},
	}
	return &service
}

func (s *SchedulerService) SendToDownload(page model.Page) {
	link := model.PageToLink(page)

	bytes, _ := json.Marshal(link)

	msg := &sarama.ProducerMessage{
		Topic: "download-links",
		Key:   sarama.StringEncoder(link.ID.String()),
		Value: sarama.ByteEncoder(bytes),
	}

	_, _, err := s.SyncProducer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return
	}
}

func (s SchedulerService) Run() {
	go func() {

	}()
}
