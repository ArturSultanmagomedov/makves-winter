package main

import (
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

const addr = "localhost:8080" // TODO: взять из конфига

func run() error {
	// Для данной задачи логика слишком простая, чтобы добавлять сервисный слой
	repo, err := NewRepository()
	if err != nil {
		return err
	}
	handler := NewHandler(repo)

	http.HandleFunc("/get-items", handler.GetItems)
	return http.ListenAndServe(addr, nil)
}
