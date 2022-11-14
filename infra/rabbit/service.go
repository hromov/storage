package rabbit

import (
	"log"

	"github.com/hromov/storage/domain/storage"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) Push(event storage.Event) error {
	log.Println("pushed event", event)
	return nil
}
