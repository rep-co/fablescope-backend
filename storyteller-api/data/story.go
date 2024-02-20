package data

// TODO: When billing and vpn issues will be fixed
// Refactor story int proper way (adding paragraphs etc.)
type Story struct {
	Content string
}

func NewStory(content string) *Story {
	return &Story{
		Content: content,
	}
}

func (s *Story) IsEmpty() bool {
	return s.Content == ""
}
