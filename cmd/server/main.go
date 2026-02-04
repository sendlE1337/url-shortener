package main

import (
	http_internal "SHORTNERED_URL/internal/http"
	shortener "SHORTNERED_URL/internal/service"
	"SHORTNERED_URL/internal/storage"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://shortener:secret@localhost:5432/shortener"
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	db := storage.NewPostgres(pool)
	//Создаем сервис, который работает с хранилишем
	service := shortener.NewService(db)

	//Создаем роутер с handler-ами для POST AND GET
	router := http_internal.NewRouter(service)

	fmt.Println("Server listening on :8080")
	//Запускаем сервер
	log.Fatal(http.ListenAndServe(":8080", router))
}
