package storygenerator

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/sashabaranov/go-openai"
)

type OpenAIStoryGenerator struct {
	client  *openai.Client
	prompt  string
	enabled bool
	mu      sync.Mutex
}

func NewOpenAIStoryGenerator(apiKey, prompt string) *OpenAIStoryGenerator {
	s := &OpenAIStoryGenerator{
		client: openai.NewClient(apiKey),
		prompt: prompt,
	}

	log.Printf("openai story generator enabled: %v", apiKey != "")

	if apiKey != "" {
		s.enabled = true
	}

	return s
}

func (s *OpenAIStoryGenerator) GenerateStory(
	ctx context.Context,
	tags []string,
) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.enabled {
		return "disabled", nil
	}

	request := openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: fmt.Sprintf("%s%s", s.prompt, strings.Join(tags, ", ")),
			},
		},
		MaxTokens:   720, //Token = 75% of a word, so here we pick story lenght
		Temperature: 0.7, //As closer to 1.0 as more creative and vice versa
		TopP:        1,   //As closer to 1.0 as more it will use natural language
	}

	resp, err := s.client.CreateChatCompletion(ctx, request)
	if err != nil {
		return "", err
	}

	//Beautify responce
	//TODO: split story into paragraphs
	rawStory := strings.TrimSpace(resp.Choices[0].Message.Content)

	return rawStory, nil
}
