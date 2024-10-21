package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writter http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(writter, "hello")
	} else {
		fmt.Fprintf(writter, "Hello "+name+" Apa kabar "+name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Eko", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameter(writter http.ResponseWriter, request *http.Request) {
	fn := request.URL.Query().Get("fn")
	ln := request.URL.Query().Get("ln")
	fmt.Fprintf(writter, "Hello "+fn+" "+ln)
}

func TestQueryMultiParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?fn=Mario&ln=Aprilnino", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(writter http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprintf(writter, "HI "+strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Mario&name=Aprilnino", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
