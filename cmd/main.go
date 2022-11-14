package main

import (
	"log"
	"net/http"

	"github.com/hromov/storage/domain/storage"
	"github.com/hromov/storage/handlers"
	"github.com/hromov/storage/infra/files"
	"github.com/hromov/storage/infra/rabbit"
)

func main() {
	fs := files.NewService()
	es := rabbit.NewService()
	ss := storage.NewService(fs, es)

	handler := handlers.Init(ss)

	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
