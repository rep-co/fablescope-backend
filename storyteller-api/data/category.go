package data

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tags []*Tag `json:"tags"`
}

type Categories struct {
	Categories []Category `json:"categories"`
}
