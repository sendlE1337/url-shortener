package main

import (
	http_internal "SHORTNERED_URL/internal/http"
	shortener "SHORTNERED_URL/internal/service"
	"SHORTNERED_URL/internal/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Создаем in-memory хранилище
	memStorage := storage.NewInMemory()

	//Создаем сервис, который работает с хранилишем
	service := shortener.NewService(memStorage)

	//Создаем роутер с handler-ами для POST AND GET
	router := http_internal.NewRouter(service)

	fmt.Println("Server listening on :8080")
	//Запускаем сервер
	log.Fatal(http.ListenAndServe(":8080", router))
}
