package db

import "github.com/RokibulHasan7/book-api/model"

var Authors []model.Author

var AuthorMap = make(map[string]model.Author)

func InitAuthor() {
	Authors = []model.Author{
		{
			FirstName: "Jokha",
			LastName:  "Alharthi",
			Email:     "jokha@gmail.com",
		},
		{
			FirstName: "Richard",
			LastName:  "Powers",
			Email:     "richar@gmail.com",
		},
		{
			FirstName: "Vasdev",
			LastName:  "Mohi",
			Email:     "mohi@gmail.com",
		},
	}
	AuthorMap := make(map[string]model.Author)
	// Mapping Author to AuthorMap; key: email, value: author
	for _, author := range Authors {
		_, ok := AuthorMap[author.Email]
		if !ok {
			AuthorMap[author.Email] = author
		}
	}
}
