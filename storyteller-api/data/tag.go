package data

import "github.com/rep-co/fablescope-backend/storyteller-api/types"

type Tag struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Icon        types.SvgString `json:"icon_svg"`
}

type TagName struct {
	Name string `json:"name"`
}
type TagNames struct {
	TagNames []TagName `json:"tags"`
}
