package models

//Films app
type Films struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	ImageURL string `json:"imageUrl"`
	Synopsis string `json:"Synopsis"`
}
