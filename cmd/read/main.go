package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-reader-blob"

	"github.com/whosonfirst/go-reader/app/read"
)

func main() {

	ctx := context.Background()
	err := read.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to read, %v", err)
	}
}
