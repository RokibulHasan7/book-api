package db

import "github.com/RokibulHasan7/book-api/model"

var Books []model.Book
var BookMap map[string]model.Book

func InitBook() {
	Books = []model.Book{
		{
			Id:          "1",
			Name:        "Celestial Bodies",
			ISBN:        "12-12-12",
			PublishDate: "01/01/2019",
			AuthorList: []model.Author{
				{
					FirstName: "Jokha",
					LastName:  "Alharthi",
					Email:     "jokha@gmail.com",
				},
				{
					FirstName: "Rokibul",
					LastName:  "Hasan",
					Email:     "rakib@gmai.com",
				},
			},
		},
		{
			Id:          "2",
			Name:        "The Overstory",
			ISBN:        "13-13-13",
			PublishDate: "02/02/2019",
			AuthorList: []model.Author{
				{
					FirstName: "Richard",
					LastName:  "Powers",
					Email:     "richar@gmail.com",
				},
			},
		},
		{
			Id:          "3",
			Name:        "Cheque book",
			ISBN:        "14-14-14",
			PublishDate: "03/03/2019",
			AuthorList: []model.Author{
				{
					FirstName: "Vasdev",
					LastName:  "Mohi",
					Email:     "mohi@gmail.com",
				},
			},
		},
	}

	// Mapping books to BookMap; key = ISBN, value = book
	for _, book := range Books {
		BookMap[book.ISBN] = book
		for _, author := range book.AuthorList {
			_, ok := AuthorMap[author.Email]
			if !ok {
				Authors = append(Authors, author)
				AuthorMap[author.Email] = author
			}
		}
	}
}
