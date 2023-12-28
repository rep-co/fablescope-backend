package data

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tags []*Tag `json:"tags"`
}

type Tag struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SvgString   string `json:"svg_string"`
}
