package storygenerator

import (
	"context"

	"github.com/rep-co/fablescope-backend/storyteller-api/data"
)

type StoryGenerator interface {
	GenerateStory(
		ctx context.Context,
		tags string,
	) (data.Story, error)
}
