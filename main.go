// static-files.go
package main

import (
	"html/template"
	"net/http"
)

func main() {
    tmpl := template.Must(template.ParseFiles("home.html", "shop.html", "about.html", "contact.html"))

    fs := http.FileServer(http.Dir("/assets/"))
    http.Handle("/fs", http.StripPrefix("/fs/", fs))
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, "home.html")
    })
    http.ListenAndServe(":8080", nil)
}
