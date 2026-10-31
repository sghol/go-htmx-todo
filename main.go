package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Task struct {
	ID     int
	Title  string
	Status string
	Done   bool
}

var tasks = []Task{
	{ID: 1, Title: "شستن صورت", Status: "active", Done: false},
	{ID: 2, Title: "خاموش کردن کامپیوتر", Status: "active", Done: false},
	{ID: 3, Title: "پرهیز از مشاجره ذهنی", Status: "active", Done: false},
}

// renderTemplate render template with data
func renderTemplate(w http.ResponseWriter, t *template.Template, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Serve htmx to the browser with correct header metadata
// handles /static/htmx.min.js
func ServeHTMX(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(w, r, "./static/htmx.min.js.gz")
}

// ServeStaticFiles css, imgates, assets, ...
func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	fs.ServeHTTP(w, r)
}

// Index home page
// handles GET /
func Index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	renderTemplate(w, t, tasks)
}

func main() {
	// handlers
	http.HandleFunc("GET /", Index)
	http.HandleFunc("GET /static/htmx.min.js", ServeHTMX)
	http.HandleFunc("GET /static/", ServeStaticFiles)

	// server
	fmt.Println("Running on port http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
