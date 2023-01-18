package model

type Book struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	ISBN        string   `json:"isbn"`
	PublishDate string   `json:"publishDate"`
	AuthorList  []Author `json:"author"`
}
