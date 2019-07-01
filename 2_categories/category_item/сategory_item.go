package category_item

//easyjson:json
type CategoryItem struct {
	Title   string         `json:"title"`
	Href    string         `json:"href"`
	Childes []CategoryItem `json:"childes,omitempty"`
}

//easyjson:json
type CategoryItems []CategoryItem
