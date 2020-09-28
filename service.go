package main

import (
	"context"

	"github.com/go-kit/kit/log"
)

type Book struct {
	BookId    string `json:"bookId,omitempty"`
	Title     string `json:"title,omitempty"`
	Edition   string `json:"edition,omitempty"`
	Copyright string `json:"copyright,omitempty"`
	Language  string `json:"language,omitempty"`
	Pages     string `json:"pages,omitempty"`
	Author    string `json:"author,omitempty"`
	Publisher string `json:"publisher,omitempty"`
}

type bookservice struct {
	logger log.Logger
}

// Service describes the Book service.
type BookService interface {
	CreateBook(ctx context.Context, book Book) (string, error)
	GetBookById(ctx context.Context, id string) (interface{}, error)
	UpdateBook(ctx context.Context, book Book) (string, error)
	DeleteBook(ctx context.Context, id string) (string, error)
}

var books = []Book{
	Book{BookId: "Book1", Title: "Operating System Concepts", Edition: "9th",
		Copyright: "2012", Language: "ENGLISH", Pages: "976",
		Author: "Abraham Silberschatz", Publisher: "John Wiley & Sons"},
	Book{BookId: "Book3", Title: "Computer Networks", Edition: "5th",
		Copyright: "2010", Language: "ENGLISH", Pages: "960",
		Author: "Andrew S. Tanenbaum", Publisher: "Andrew S. Tanenbaum"},
}

func find(x string) int {
	for i, book := range books {
		if x == book.BookId {
			return i
		}
	}
	return -1
}

func NewService(logger log.Logger) BookService {
	return &bookservice{
		logger: logger,
	}
}

func (s bookservice) CreateBook(ctx context.Context, book Book) (string, error) {
	var msg = "success"
	books = append(books, book)
	return msg, nil
}

func (s bookservice) GetBookById(ctx context.Context, id string) (interface{}, error) {
	var err error
	var book interface{}
	var empty interface{}
	i := find(id)
	if i == -1 {
		return empty, err
	}
	book = books[i]
	return book, nil
}
func (s bookservice) DeleteBook(ctx context.Context, id string) (string, error) {
	var err error
	msg := ""
	i := find(id)
	if i == -1 {
		return "", err
	}
	copy(books[i:], books[i+1:])
	books[len(books)-1] = Book{}
	books = books[:len(books)-1]
	return msg, nil
}
func (s bookservice) UpdateBook(ctx context.Context, book Book) (string, error) {
	var empty = ""
	var err error
	var msg = "success"
	i := find(book.BookId)
	if i == -1 {
		return empty, err
	}
	books[i] = book
	return msg, nil
}

type Author struct {
	AuthorId    string `json:"authorId,omitempty"`
	Name        string `json:"name,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	Birth       string `json:"birth,omitempty"`
	Genere      string `json:"genere,omitempty"`
}

type authorservice struct {
	logger log.Logger
}

// Service describes the author service.
type AuthorService interface {
	CreateAuthor(ctx context.Context, author Author) (string, error)
	GetAuthorById(ctx context.Context, id string) (interface{}, error)
	UpdateAuthor(ctx context.Context, author Author) (string, error)
	DeleteAuthor(ctx context.Context, id string) (string, error)
}

var authors = []Author{
	Author{AuthorId: "Author1", Name: "Abraham Silberschatz", Nationality: "English",
		Birth: "12/05/1973", Genere: "Computer science engineering"},
	Author{AuthorId: "Author2", Name: "Andrew S. Tanenbaum", Nationality: "English",
		Birth: "05/06/1965", Genere: "Computer science engineering"},
}

func findAuthor(x string) int {
	for i, author := range authors {
		if x == author.AuthorId {
			return i
		}
	}
	return -1
}

func NewService_Author(logger log.Logger) AuthorService {
	return &authorservice{
		logger: logger,
	}
}

func (s authorservice) CreateAuthor(ctx context.Context, author Author) (string, error) {
	var msg = "success"
	authors = append(authors, author)
	return msg, nil
}

func (s authorservice) GetAuthorById(ctx context.Context, id string) (interface{}, error) {
	var err error
	var author interface{}
	var empty interface{}
	i := findAuthor(id)
	if i == -1 {
		return empty, err
	}
	author = authors[i]
	return author, nil
}

func (s authorservice) DeleteAuthor(ctx context.Context, id string) (string, error) {
	var err error
	msg := ""
	i := findAuthor(id)
	if i == -1 {
		return "", err
	}
	copy(authors[i:], authors[i+1:])
	authors[len(authors)-1] = Author{}
	authors = authors[:len(authors)-1]
	return msg, nil
}

func (s authorservice) UpdateAuthor(ctx context.Context, author Author) (string, error) {
	var empty = ""
	var err error
	var msg = "success"
	i := findAuthor(author.AuthorId)
	if i == -1 {
		return empty, err
	}
	authors[i] = author
	return msg, nil
}

type Publisher struct {
	PublisherId string `json:"publisherId,omitempty"`
	Name        string `json:"name,omitempty"`
	Country     string `json:"country,omitempty"`
	Founded     string `json:"founded,omitempty"`
	Genere      string `json:"genere,omitempty"`
}

type publisherservice struct {
	logger log.Logger
}

// Service describes the author service.
type PublisherService interface {
	CreatePublisher(ctx context.Context, publisher Publisher) (string, error)
	GetPublisherById(ctx context.Context, id string) (interface{}, error)
	UpdatePublisher(ctx context.Context, publisher Publisher) (string, error)
	DeletePublisher(ctx context.Context, id string) (string, error)
}

var publishers = []Publisher{
	Publisher{PublisherId: "Publisher1", Name: "John Wiley & Sons", Country: "England",
		Founded: "05/07/1945", Genere: "Computer science engineering"},
	Publisher{PublisherId: "Publisher2", Name: "Andrew S. Tanenbaum", Country: "England",
		Founded: "08/09/1965", Genere: "Computer science engineering"},
}

func findPublisher(x string) int {
	for i, publisher := range publishers {
		if x == publisher.PublisherId {
			return i
		}
	}
	return -1
}

func NewService_Publisher(logger log.Logger) PublisherService {
	return &publisherservice{
		logger: logger,
	}
}

func (s publisherservice) CreatePublisher(ctx context.Context, publisher Publisher) (string, error) {
	var msg = "success"
	publishers = append(publishers, publisher)
	return msg, nil
}

func (s publisherservice) GetPublisherById(ctx context.Context, id string) (interface{}, error) {
	var err error
	var publisher interface{}
	var empty interface{}
	i := findPublisher(id)
	if i == -1 {
		return empty, err
	}
	publisher = publishers[i]
	return publisher, nil
}

func (s publisherservice) DeletePublisher(ctx context.Context, id string) (string, error) {
	var err error
	msg := ""
	i := findPublisher(id)
	if i == -1 {
		return "", err
	}

	copy(publishers[i:], publishers[i+1:])
	publishers[len(publishers)-1] = Publisher{}
	publishers = publishers[:len(publishers)-1]
	return msg, nil
}

func (s publisherservice) UpdatePublisher(ctx context.Context, publisher Publisher) (string, error) {
	var empty = ""
	var err error
	var msg = "success"
	i := findPublisher(publisher.PublisherId)
	if i == -1 {
		return empty, err
	}
	publishers[i] = publisher
	return msg, nil
}

type PublishersBooks struct {
	PublisherId string `json:"publisherId,omitempty"`
	BookId      string `json:"bookId,omitempty"`
}

type publishersBooksservice struct {
	logger log.Logger
}

// Service describes the author service.
type PublishersBooksService interface {
	GetPublisherByBook(ctx context.Context, id string) (interface{}, error)
	GetBooksByPublisher(ctx context.Context, id string) (interface{}, error)
}

var publishers_books = []PublishersBooks{
	PublishersBooks{PublisherId: "Publisher1", BookId: "Book1"},
	PublishersBooks{PublisherId: "Publisher2", BookId: "Book3"},
}

func findPublisherByBook(x string) int {
	for i, publishers_books := range publishers_books {
		if x == publishers_books.BookId {
			return i
		}
	}
	return -1
}

func findBooksByPublisher(x string) int {
	for i, publishers_books := range publishers_books {
		if x == publishers_books.PublisherId {
			return i
		}
	}
	return -1
}

func (s publishersBooksservice) GetPublisherByBook(ctx context.Context, id string) (interface{}, error) {
	var err error
	var publisher interface{}
	var empty interface{}
	i := findPublisherByBook(id)
	if i == -1 {
		return empty, err
	}
	publisher = publishers[i]
	return publisher, nil
}

func (s publishersBooksservice) GetBooksByPublisher(ctx context.Context, id string) (interface{}, error) {
	var err error
	var book interface{}
	var empty interface{}
	i := findBooksByPublisher(id)
	if i == -1 {
		return empty, err
	}
	book = books[i]
	return book, nil
}

func NewService_PublishersBooks(logger log.Logger) PublishersBooksService {
	return &publishersBooksservice{
		logger: logger,
	}
}

type AuthorsBooks struct {
	AuthorId string `json:"authorId,omitempty"`
	BookId   string `json:"bookId,omitempty"`
}

type authorsBooksservice struct {
	logger log.Logger
}

// Service describes the author service.
type AuthorsBooksService interface {
	GetAuthorByBook(ctx context.Context, id string) (interface{}, error)
	GetBooksByAuthor(ctx context.Context, id string) (interface{}, error)
}

var authors_books = []AuthorsBooks{
	AuthorsBooks{AuthorId: "Author1", BookId: "Book1"},
	AuthorsBooks{AuthorId: "Author2", BookId: "Book3"},
}

func finAuthorByBook(x string) int {
	for i, authors_books := range authors_books {
		if x == authors_books.BookId {
			return i
		}
	}
	return -1
}

func findBooksByAuthors(x string) int {
	for i, authors_books := range authors_books {
		if x == authors_books.AuthorId {
			return i
		}
	}
	return -1
}

func (s authorsBooksservice) GetAuthorByBook(ctx context.Context, id string) (interface{}, error) {
	var err error
	var author interface{}
	var empty interface{}
	i := finAuthorByBook(id)
	if i == -1 {
		return empty, err
	}
	author = authors[i]
	return author, nil
}

func (s authorsBooksservice) GetBooksByAuthor(ctx context.Context, id string) (interface{}, error) {
	var err error
	var book interface{}
	var empty interface{}
	i := findBooksByAuthors(id)
	if i == -1 {
		return empty, err
	}
	book = books[i]
	return book, nil
}

func NewService_AuthorsBooks(logger log.Logger) AuthorsBooksService {
	return &authorsBooksservice{
		logger: logger,
	}
}
