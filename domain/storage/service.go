package storage

import (
	"io"

	"github.com/google/uuid"
)

//go:generate mockgen -source=service.go -destination=../../mocks/mock_storage_service.go -package=mocks FileService
type FileService interface {
	Save(c FileCategory, fileName string, file io.Reader) error
}

//go:generate mockgen -source=service.go -destination=../../mocks/mock_storage_service.go -package=mocks EventsService
type EventsService interface {
	Push(Event) error
}

type service struct {
	fs FileService
	es EventsService
}

func NewService(fs FileService, es EventsService) *service {
	return &service{fs, es}
}

func (s *service) Save(c FileCategory, file io.Reader) error {
	fileName := uuid.New().String()
	if err := s.fs.Save(c, fileName, file); err != nil {
		return err
	}

	event := Event{
		Category: c,
		Type:     EventCreate,
		FileName: fileName,
	}

	if err := s.es.Push(event); err != nil {
		return err
	}
	//TODO: retry ? or just log error?
	return nil
}
