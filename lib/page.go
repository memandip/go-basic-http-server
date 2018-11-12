package page

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))
var validPath = regexp.MustCompile("^/(view|save|edit|/)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

func GetTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	defer CatchError()

	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page title")
	}
	return m[2], nil
}

func (p *Page) Save() error {
	defer CatchError()

	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	defer CatchError()

	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{title, body}, nil
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, title string, pageType string) {
	defer CatchError()
	p, err := LoadPage(title)

	if err != nil {
		page := Page{Title: title}
		page.Save()
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}

	templatePath := "templates/" + pageType + ".html"
	// error := templates.ExecuteTemplate(w, templatePath, p)
	// fmt.Println("error", error)
	// if error != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	t, _ := template.ParseFiles(templatePath)
	t.Execute(w, p)
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func CatchError() {
	if r := recover(); r != nil {
		fmt.Println("Error message", r)
		// panic("Something went wrong")
	}
}
