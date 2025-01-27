package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googleai"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if err := googleai.Init(ctx, nil); err != nil {
		log.Fatal("Error initializing googleai: ", err)
	}

	genkit.DefineFlow("menuSuggestionFlow", func(ctx context.Context, input string) (string, error) {
		m := googleai.Model("gemini-1.5-flash")
		if m == nil {
			return "", errors.New("menuSuggestionFlow: failed to find model")
		}

		resp, err := ai.Generate(ctx, m,
			ai.WithConfig(&ai.GenerationCommonConfig{Temperature: 1}),
			ai.WithTextPrompt(fmt.Sprintf(`Suggest an item for the menu of a %s themed restaurant`, input)),
		)
		if err != nil {
			return "", err
		}

		text := resp.Text()
		return text, nil
	})

	if err := genkit.Init(ctx, nil); err != nil {
		log.Fatal("Error initializing genkit: ", err)
	}
}
