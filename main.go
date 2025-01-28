package main

import (
	"context"
	"go-genkit-googleai/config"
	"log"
	"os"

	"github.com/firebase/genkit/go/genkit"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	plugins := os.Getenv("PLUGINS")
	if plugins == "googleai" {
		config.GoogleAI(ctx)
	} else if plugins == "vertexai" {
		config.VertexAI(ctx)
	} else {
		log.Fatal("Unknown plugin")
	}

	if err := genkit.Init(ctx, nil); err != nil {
		log.Fatal("Error initializing genkit: ", err)
	}
}
