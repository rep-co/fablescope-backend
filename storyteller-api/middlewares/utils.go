package middlewares

var (
	contextKeyTags  = contextKey("tags")
	contextKeyStory = contextKey("story")
)

type contextKey string

func (c contextKey) String() string {
	return "middlewares context key " + string(c)
}
