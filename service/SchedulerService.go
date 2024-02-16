package service

import (
	"encoding/json"
	"log"
	"mongo_peg_1/model"
	"mongo_peg_1/repository"

	"github.com/IBM/sarama"
	"github.com/hashicorp/go-uuid"
)

type SchedulerService struct {
	SyncProducer   sarama.SyncProducer
	PageRepository repository.PageRepository
	Links          []model.Link
}

func NewSchedulerService() *SchedulerService {
	var service = SchedulerService{}
	service.Links = []model.Link{
		{Url: "", Status: ""},
		{Url: "", Status: ""},
		{Url: "", Status: ""},
		{Url: "", Status: ""},
		{Url: "", Status: ""},
	}
	return &service
}

func (s *SchedulerService) SendToDownload(Pages []model.Page) {
	id, _ := uuid.GenerateUUID()
	message := model.Link{
		Id:  id,
		Url: "",
	}

	// Преобразование сообщения в JSON что бы потом отправить через kafka
	bytes, _ := json.Marshal(message)

	msg := &sarama.ProducerMessage{
		Topic: "ping",
		Key:   sarama.StringEncoder(id),
		Value: sarama.ByteEncoder(bytes),
	}

	// отправка сообщения в Kafka
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
