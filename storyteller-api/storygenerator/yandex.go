package storygenerator

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/rep-co/fablescope-backend/storyteller-api/data"
	"github.com/sheeiavellie/go-yandexgpt"
)

type YandexStoryGenerator struct {
	client    *yandexgpt.YandexGPTClient
	prompt    string
	catalogID string
	enabled   bool
	mu        sync.Mutex
}

func NewYandexStoryGeneratorWithAPIKey(
	apiKey,
	catalogID,
	prompt string,
) *YandexStoryGenerator {
	s := &YandexStoryGenerator{
		client:    yandexgpt.NewYandexGPTClientWithAPIKey(apiKey),
		prompt:    prompt,
		catalogID: catalogID,
	}
	log.Printf("yandexgpt story generator enabled: %v", apiKey != "")

	if apiKey != "" {
		s.enabled = true
	}

	return s
}

func NewYandexStoryGeneratorWithIAMToken(
	iamToken,
	catalogID,
	prompt string,
) *YandexStoryGenerator {
	s := &YandexStoryGenerator{
		client:    yandexgpt.NewYandexGPTClientWithIAMToken(iamToken),
		prompt:    prompt,
		catalogID: catalogID,
	}
	log.Printf("yandexgpt story generator enabled: %v", iamToken != "")

	if iamToken != "" {
		s.enabled = true
	}

	return s
}

func (s *YandexStoryGenerator) GenerateStory(
	ctx context.Context,
	tagNames []data.TagName,
) (*data.Story, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.enabled {
		return data.NewStoryEmpty(), fmt.Errorf("story generator is not enabled")
	}

	tags := tagNamesToString(tagNames)

	request := yandexgpt.YandexGPTRequest{
		ModelURI: yandexgpt.MakeModelURI(s.catalogID, yandexgpt.YandexGPTModelLite),
		CompletionOptions: yandexgpt.YandexGPTCompletionOptions{
			Stream:      false,
			Temperature: 0.7,
			MaxTokens:   2000,
		},
		Messages: []yandexgpt.YandexGPTMessage{
			{
				Role: yandexgpt.YandexGPTMessageRoleSystem,
				Text: s.prompt,
			},
			{
				Role: yandexgpt.YandexGPTMessageRoleUser,
				Text: fmt.Sprintf("Теги: %s", tags),
			},
		},
	}
	response, err := s.client.CreateRequest(ctx, request)
	if err != nil {
		return data.NewStoryEmpty(), err
	}

	rawStory := response.Result.Alternatives[0].Message.Text
	return data.NewStory(rawStory), nil
}
