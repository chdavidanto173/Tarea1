package main

import (
	"context"

	pb "github.com/chdavidanto173/Tarea1/booksapp"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	bookMap map[string]*pb.Book
}

func (s *server) AddBook(ctx context.Context, in *pb.Book) (*pb.BookID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating Book ID", err)
	}
	in.Id = out.String()
	if s.bookMap == nil {
		s.bookMap = make(map[string]*pb.Book)
	}
	s.bookMap[in.Id] = in
	return &pb.BookID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) DeleteBook(ctx context.Context, in *pb.BookID) (*pb.Book, error) {
	value, exists := s.bookMap[in.Value]
	if exists {
		delete(s.bookMap, in.Value)
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Book does not exist.", in.Value)
}

func (s *server) GetBook(ctx context.Context, in *pb.BookID) (*pb.Book, error) {
	value, exists := s.bookMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Book does not exist.", in.Value)
}

func (s *server) UpdateBook(ctx context.Context, in *pb.Book) (*pb.Book, error) {
	value, exists := s.bookMap[in.Id]
	if exists {

		if in.Language != "" {
			value.Language = in.Language
		}

		if in.Copyright != "" {
			value.Copyright = in.Copyright
		}

		if in.Edition != "" {
			value.Edition = in.Edition
		}

		if in.Pages != "" {
			value.Pages = in.Pages
		}

		if in.Title != "" {
			value.Title = in.Title
		}

		if in.Publisher != "" {
			value.Publisher = in.Publisher
		}

		if in.Author != "" {
			value.Author = in.Author
		}

		s.bookMap[in.Id] = value
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Book does not exist.", in.Id)
}
