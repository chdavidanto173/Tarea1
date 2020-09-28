package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	/*
			logger := log.NewLogfmtLogger(os.Stderr)

			r := mux.NewRouter()

			var svc BookService
			svc = NewService(logger)

			// svc = loggingMiddleware{logger, svc}
			// svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

			CreateBookHandler := httptransport.NewServer(
				makeCreateBookEndpoint(svc),
				decodeCreateBookRequest,
				encodeResponse,
			)
			GetByBookIdHandler := httptransport.NewServer(
				makeGetBookByIdEndpoint(svc),
				decodeGetBookByIdRequest,
				encodeResponse,
			)
			DeleteBookHandler := httptransport.NewServer(
				makeDeleteBookEndpoint(svc),
				decodeDeleteBookRequest,
				encodeResponse,
			)
			UpdateBookHandler := httptransport.NewServer(
				makeUpdateBookendpoint(svc),
				decodeUpdateBookRequest,
				encodeResponse,
			)
			http.Handle("/", r)
			http.Handle("/book", CreateBookHandler)
			http.Handle("/book/update", UpdateBookHandler)
			r.Handle("/book/{bookid}", GetByBookIdHandler).Methods("GET")
			r.Handle("/book/{bookid}", DeleteBookHandler).Methods("DELETE")

			// http.Handle("/metrics", promhttp.Handler())
			logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
			logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))

			logger2 := log.NewLogfmtLogger(os.Stderr)

			r2 := mux.NewRouter()

			var svcA AuthorService
			svcA = NewService_Author(logger2)

			CreateAuthorHandler := httptransport.NewServer(
				makeCreateAuthorEndpoint(svcA),
				decodeCreateAuthorRequest,
				encodeResponseAuthor,
			)
			GetByAuthorIdHandler := httptransport.NewServer(
				makeGetAuthorByIdEndpoint(svcA),
				decodeGetAuthorByIdRequest,
				encodeResponseAuthor,
			)
			DeleteAuthorHandler := httptransport.NewServer(
				makeDeleteAuthorEndpoint(svcA),
				decodeDeleteAuthorRequest,
				encodeResponseAuthor,
			)
			UpdateAuthorHandler := httptransport.NewServer(
				makeUpdateAuthorendpoint(svcA),
				decodeUpdateAuthorRequest,
				encodeResponseAuthor,
			)

			http.Handle("/", r2)
			http.Handle("/author", CreateAuthorHandler)
			http.Handle("/author/update", UpdateAuthorHandler)
			r2.Handle("/author/{authorid}", GetByAuthorIdHandler).Methods("GET")
			r2.Handle("/author/{authorid}", DeleteAuthorHandler).Methods("DELETE")

			// http.Handle("/metrics", promhttp.Handler())
			logger2.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
			logger2.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))

			logger3 := log.NewLogfmtLogger(os.Stderr)

			r3 := mux.NewRouter()

			var svcP PublisherService
			svcP = NewService_Publisher(logger3)

			CreatePublisherHandler := httptransport.NewServer(
				makeCreatePublisherEndpoint(svcP),
				decodeCreatePublisherRequest,
				encodeResponsePublisher,
			)
			GetByPublisherIdHandler := httptransport.NewServer(
				makeGetPublisherByIdEndpoint(svcP),
				decodeGetPublisherByIdRequest,
				encodeResponsePublisher,
			)
			DeletePublisherHandler := httptransport.NewServer(
				makeDeletePublisherEndpoint(svcP),
				decodeDeletePublisherRequest,
				encodeResponsePublisher,
			)
			UpdatePublisherHandler := httptransport.NewServer(
				makeUpdatePublisherendpoint(svcP),
				decodeUpdatePublisherRequest,
				encodeResponsePublisher,
			)

			http.Handle("/", r3)
			http.Handle("/publisher", CreatePublisherHandler)
			http.Handle("/publisher/update", UpdatePublisherHandler)
			r3.Handle("/publisher/{publisherid}", GetByPublisherIdHandler).Methods("GET")
			r3.Handle("/publisher/{publisherid}", DeletePublisherHandler).Methods("DELETE")

			// http.Handle("/metrics", promhttp.Handler())
			logger3.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
			logger3.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))


		logger4 := log.NewLogfmtLogger(os.Stderr)

		r4 := mux.NewRouter()

		var svcPB PublishersBooksService
		svcPB = NewService_PublishersBooks(logger4)

		GetByPublisherByBookHandler := httptransport.NewServer(
			makeGetPublisherByBookEndpoint(svcPB),
			decodeGetPublisherByBookRequest,
			encodeResponsePublishersBook,
		)

		GetBooksByPublisherHandler := httptransport.NewServer(
			makeGetBooksByPublisherEndpoint(svcPB),
			decodeGetBooksByPublisherRequest,
			encodeResponsePublishersBook,
		)

		http.Handle("/", r4)
		r4.Handle("/books/{bookid}/publishers", GetByPublisherByBookHandler).Methods("GET")
		r4.Handle("/publishers/{publisherid}/books", GetBooksByPublisherHandler).Methods("GET")

		logger4.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
		logger4.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))

	*/

	logger5 := log.NewLogfmtLogger(os.Stderr)

	r5 := mux.NewRouter()

	var svcPB AuthorsBooksService
	svcPB = NewService_AuthorsBooks(logger5)

	GetAuthorByBookHandler := httptransport.NewServer(
		makeGetAuthorByBookEndpoint(svcPB),
		decodeGetAuthorByBookRequest,
		encodeResponseAuthorsBook,
	)

	GetBooksByAuthorHandler := httptransport.NewServer(
		makeGetBooksByAuthorEndpoint(svcPB),
		decodeGetBooksByAuthorRequest,
		encodeResponseAuthorsBook,
	)

	http.Handle("/", r5)
	r5.Handle("/books/{bookid}/author", GetAuthorByBookHandler).Methods("GET")
	r5.Handle("/authors/{authorid}/books", GetBooksByAuthorHandler).Methods("GET")

	logger5.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
	logger5.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
