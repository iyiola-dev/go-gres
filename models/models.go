package models

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:Author`
	Year   string `json:Year`
}

var books []Book
