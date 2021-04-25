package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/iyiola-dev/go-gres/models"
)

var book []models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	bookslice := append(book,
		models.Book{ID: 1, Title: "montenegro", Author: "name", Year: 2020},
		models.Book{ID: 2, Title: "chiriyeo", Author: "jon", Year: 1999},
		models.Book{ID: 3, Title: "mistl", Author: "varys", Year: 1974},
		models.Book{ID: 4, Title: "chantant", Author: "lutehn", Year: 1938},
		models.Book{ID: 5, Title: "vantahen", Author: "vueh", Year: 2016})

	json.NewEncoder(w).Encode(bookslice)

}
func GetBook(w http.ResponseWriter, r *http.Request) {
	bookslice := slic()
	params := mux.Vars(r)
	log.Println(params)
	i, _ := strconv.Atoi(params["id"])
	for _, books := range bookslice {
		if books.ID == i {
			json.NewEncoder(w).Encode(&books)

		}
	}

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var bookmodel models.Book
	json.NewDecoder(r.Body).Decode(&bookmodel)
	bookslice := slic()

	for i, items := range bookslice {
		if items.ID == bookmodel.ID {
			bookslice[i] = bookmodel
			log.Println(bookslice[i].ID)
		}
	}
	json.NewEncoder(w).Encode(bookslice)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params["id"])
	id, _ := strconv.Atoi(params["id"])
	booklice := slic()
	for i, items := range booklice {
		if items.ID == id {
			books := append(booklice[:i], booklice[i+1:]...)
			json.NewEncoder(w).Encode(books)
		}
	}
}
func AddBook(w http.ResponseWriter, r *http.Request) {
	bookslice := slic()
	var bookmodel models.Book
	json.NewDecoder(r.Body).Decode(&bookmodel)
	block := append(bookslice, bookmodel)
	json.NewEncoder(w).Encode(block)

}

func slic() []models.Book {
	bookslice := append(book,
		models.Book{ID: 1, Title: "montenegro", Author: "name", Year: 2020},
		models.Book{ID: 2, Title: "chiriyeo", Author: "jon", Year: 1999},
		models.Book{ID: 3, Title: "mistl", Author: "varys", Year: 1974},
		models.Book{ID: 4, Title: "chantant", Author: "lutehn", Year: 1938},
		models.Book{ID: 5, Title: "vantahen", Author: "vueh", Year: 2016})
	return bookslice
}
