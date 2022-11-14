package main

import (
	"log"
	"net/http"

	"github.com/hromov/storage/api/handlers"
	"github.com/hromov/storage/config"
	"github.com/hromov/storage/domain/storage"
	"github.com/hromov/storage/infra/files"
	"github.com/hromov/storage/infra/rabbit"
)

func main() {
	cfg := config.Get()

	fs := files.NewService()
	es := rabbit.NewService()
	ss := storage.NewService(fs, es)

	handler := handlers.Init(ss)

	log.Println("server started at ", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, handler))
}
