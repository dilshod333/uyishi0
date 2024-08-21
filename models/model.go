package models


type Books struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	PublishedDate string `json:"published_date"`
	Isbn int64 `json:"isbn"`

}
