package main

import (
  "html/template"
  "net/http"
  "path/filepath"
  "log"
  "regexp"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))
var validPath = regexp.MustCompile("^/(|about|work|contact)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
      fn(w, r)
      return
    }

    m := validPath.FindStringSubmatch(r.URL.Path)

    if m == nil {
      notFoundHandler(w, r)
      return
    }

    fn(w, r)
  }
}

func renderTemplate(w http.ResponseWriter, t string) {
  lp := filepath.Join("templates", "layout.html")
  fp := filepath.Join("templates", t+".html")

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  tmpl.ExecuteTemplate(w, "layout", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  renderTemplate(w, "index")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
  renderTemplate(w, "about")
}

func workHandler(w http.ResponseWriter, r *http.Request) {
  renderTemplate(w, "work")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
  renderTemplate(w, "contact")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
  renderTemplate(w, "notFound")
}

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.HandleFunc("/", makeHandler(indexHandler))
  http.HandleFunc("/about", makeHandler(aboutHandler))
  http.HandleFunc("/work", makeHandler(workHandler))
  http.HandleFunc("/contact", makeHandler(contactHandler))
  log.Println("Listening...")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
