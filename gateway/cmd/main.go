package main

import (
	"context"
	"log"
	"os"

	"MarketPlace/gateway/internal/repository"
)

func main() {

	storage, err := repository.NewDbPostgres(context.Background())
	if err != nil {
		log.Printf("Fail to connect DB: %v\n", err)
		os.Exit(1)
	}
	defer storage.Close()
	log.Println("connected to auth DB")
}
