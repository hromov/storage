package files

import (
	"io"
	"log"

	"github.com/hromov/storage/domain/storage"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) Save(c storage.FileCategory, fileName string, file io.Reader) error {
	log.Println("save file", fileName, "to category", c)
	return nil
}
