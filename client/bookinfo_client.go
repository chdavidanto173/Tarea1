package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"time"

	pb "github.com/chdavidanto173/Tarea1/booksapp"
	"google.golang.org/grpc"
)

type Book struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Edition   string `json:"edition"`
	Copyright string `json:"copyright"`
	Language  string `json:"language"`
	Pages     string `json:"pages"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

var books []Book

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func readCSV(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	books = []Book{}

	for _, record := range records {
		book := Book{
			Id:        record[0],
			Title:     record[1],
			Edition:   record[2],
			Copyright: record[3],
			Language:  record[4],
			Pages:     record[5],
			Author:    record[6],
			Publisher: record[7]}
		books = append(books, book)
	}
	file.Close()
}

func main() {
	address := os.Getenv("ADDRESS")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBookInfoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddBook(ctx, &pb.Book{
		Id:        "35",
		Title:     "Big Data",
		Edition:   "11th",
		Copyright: "2010",
		Language:  "ENGLISH",
		Pages:     "975",
		Author:    "Abraham Silberschatz",
		Publisher: "John Wiley & Sons"})
	if err != nil {
		log.Fatalf("Could not add book: %v", err)
	}

	log.Printf("Book ID: %s added successfully", r.Value)

	//Ejericio 2

	book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get book: %v", err)
	}
	log.Printf("Book: ", book.String())

	book.Copyright = "2021"
	book.Language = "Spanish"
	book.Pages = "1000"

	u, erru := c.UpdateBook(ctx, book)
	if erru != nil {
		log.Fatalf("Could not update book: %v", erru)
	}

	log.Printf("Book ID: %s updated successfully", u.Id)
	log.Printf("Book: ", book.String())

	book1, err2 := c.DeleteBook(ctx, &pb.BookID{Value: r.Value})
	if err2 != nil {
		log.Fatalf("Could not get book: %v", err2)
	}
	log.Printf("Deleted Book: ", book1.String())

	//Ejercicio 3

	readCSV("books.csv")

	for _, book := range books {
		it, err := c.AddBook(ctx, &pb.Book{
			Id:        book.Id,
			Title:     book.Title,
			Edition:   book.Edition,
			Copyright: book.Copyright,
			Language:  book.Language,
			Pages:     book.Pages,
			Author:    book.Author,
			Publisher: book.Publisher})
		if err != nil {
			log.Fatalf("Could not add book: %v", err)
		} else {
			b1, _ := c.GetBook(ctx, &pb.BookID{Value: it.Value})
			log.Printf("Book: ", b1.String())
		}
	}

}
