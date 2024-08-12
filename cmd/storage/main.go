package main

import (
	"fmt"
	"log"

	"github.com/roolee10/storage/internal/storage"
)

func main() {
	storage := storage.NewStorage()

	file, err := storage.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it uploaded", file)
}
