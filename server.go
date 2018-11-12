package main

import (
	"fmt"
	"go-web/lib"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	defer page.CatchError()
	file, _ := os.Open("./data")
	filenames, _ := file.Readdirnames(0) // 0 to read all files and folders

	// t, _ := template.ParseFiles("templates/index.html")
	// data := map[string]interface{}{
	// 	"files": filenames,
	// }
	// t.Execute(w, data)
	// fmt.Println(t)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if len(filenames) == 0 {
		fmt.Fprintf(w, "No files available.\n")
		fmt.Fprintf(w, "<a href='/edit/test'>Create new</a>")
		return
	}

	template := "<h1>List of available files</h1><ul>"
	for _, name := range filenames {
		ext := filepath.Ext(name)
		template += "<li><a href='/view/" + name[:len(name)-len(ext)] + "'>" + name + "</a></li>"
	}
	template += "</ul>"
	fmt.Fprint(w, template)
	return
	// fmt.Fprintf(w, "Hi there, Path = %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page.RenderTemplate(w, r, title, "view")
	return
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	page.RenderTemplate(w, r, title, "edit")
	return
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &page.Page{Title: title, Body: []byte(body)}
	p.Save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	defer page.CatchError()

	http.HandleFunc("/view/", page.MakeHandler(viewHandler))
	http.HandleFunc("/edit/", page.MakeHandler(editHandler))
	http.HandleFunc("/save/", page.MakeHandler(saveHandler))
	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
