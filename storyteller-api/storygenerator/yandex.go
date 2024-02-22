package storygenerator

import (
	"context"
	"sync"

	"github.com/rep-co/fablescope-backend/storyteller-api/data"
)

type YandexStoryGenerator struct {
	client  string
	prompt  string
	enabled bool
	mu      sync.Mutex
}

func NewYandexStoryGenerator(
	catalogID,
	apiKey,
	prompt string,
) *YandexStoryGenerator {
	return &YandexStoryGenerator{}
}
func (s *YandexStoryGenerator) GenerateStory(
	ctx context.Context,
	tags string,
) (data.Story, error) {
	return data.Story{}, nil
}
