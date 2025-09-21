package main

import (
	"log"

	"github.com/Julia-Marcal/packet-scope/internal/application"
)

func main() {
	if err := application.StartAnalysis(); err != nil {
		log.Fatalf("Failed to start packet analysis: %v", err)
	}
}
