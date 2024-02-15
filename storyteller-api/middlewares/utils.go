package middlewares

var (
	contextKeyTags = contextKey("tags")
)

type contextKey string

func (c contextKey) String() string {
	return "middlewares context key " + string(c)
}
