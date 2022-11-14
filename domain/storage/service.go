package storage

import (
	"io"

	"github.com/google/uuid"
)

//go:generate mockgen -source=service.go -destination=../../mocks/mock_storage_service.go -package=mocks FileService
type FileService interface {
	Save(c FileCategory, fileName string, file io.Reader) error
}

//go:generate mockgen -source=service.go -destination=../../mocks/mock_storage_service.go -package=mocks NotifierService
type NotifierService interface {
	Notify(Event) error
}

type service struct {
	fs FileService
	ns NotifierService
}

func NewService(fs FileService, ns NotifierService) *service {
	return &service{fs, ns}
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

	if err := s.ns.Notify(event); err != nil {
		return err
	}
	//TODO: retry ? or just log error?
	return nil
}
