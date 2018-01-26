package lessons

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Page is an interface
type Page interface {
	Name() string
	Title() string
}

type indexPage struct {
	name  string
	title string
}

// Name is the name
func (ip indexPage) Name() string {
	return ip.name
}

// Title is the title
func (ip indexPage) Title() string {
	return ip.title
}

func indexHandle(w http.ResponseWriter, r *http.Request, p Page) {
	io.WriteString(w, fmt.Sprintf("Hello %s %s\n", p.Title(), p.Name()))
}

// RunHTTPServer runs the http server
func RunHTTPServer() {
	http.HandleFunc("/", note(reallyNote(final(indexHandle, &indexPage{
		name:  "Robert",
		title: "Mr.",
	}))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func reallyNote(h func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Really noted")
		h(w, r)
	}
}

func note(h func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Noted")
		h(w, r)
	}
}

func final(fin func(http.ResponseWriter, *http.Request, Page), p Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fin(w, r, p)
	}
}
