package config

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/vertexai"
)

func VertexAI(ctx context.Context) {
	if err := vertexai.Init(ctx, nil); err != nil {
		log.Fatal("Error initializing googleai: ", err)
	}

	genkit.DefineFlow("menuSuggestionFlow", func(ctx context.Context, input string) (string, error) {
		m := vertexai.Model("gemini-1.5-flash")
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
}
