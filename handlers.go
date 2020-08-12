package main

import (
	"encoding/json"
	"net/http"
	"path"
)

func find(x string) int {
	for i, book := range books {
		if x == book.Id {
			return i
		}
	}
	return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	if id == "book" {
		dataJson, _ := json.Marshal(books)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)

		return
	} else {

		i := find(id)
		if i == -1 {
			return
		}
		dataJson, _ := json.Marshal(books[i])
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	book := Book{}
	json.Unmarshal(body, &book)
	books = append(books, book)
	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	book := Book{}
	json.Unmarshal(body, &book)

	i := find(id)
	if i == -1 {
		return
	}

	book2 := books[i]

	if book.Title != "" {
		book2.Title = book.Title
	}

	if book.Edition != "" {
		book2.Edition = book.Edition
	}

	if book.Copyright != "" {
		book2.Copyright = book.Copyright
	}

	if book.Language != "" {
		book2.Language = book.Language
	}

	if book.Pages != "" {
		book2.Pages = book.Pages
	}

	if book.Author != "" {
		book2.Author = book.Author
	}

	if book.Publisher != "" {
		book2.Publisher = book.Publisher
	}

	books[i] = book2
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	}
	RemoveBook(books, i)
	w.WriteHeader(200)
	return
}

func RemoveBook(array []Book, i int) []Book {
	array[i] = array[len(array)-1]
	return array[:len(array)-1]

}
