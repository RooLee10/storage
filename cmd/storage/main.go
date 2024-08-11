package main

import (
	"fmt"

	"github.com/roolee10/storage/internal/file"
	"github.com/roolee10/storage/internal/storage"
)

func main() {
	storage := storage.NewStorage()
	file := file.NewFile()
	fmt.Println("it work", storage, file)
}
