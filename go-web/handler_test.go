package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
		// fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Hello World")
	})
	mux.HandleFunc("/hi", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "hi")
	})
	mux.HandleFunc("/helo", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "helo")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.URL.Path)
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
