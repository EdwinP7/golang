package main

import (
	"html/template"
	"log"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	http.ServeFile(res, req, "index.html")
	tpl.Execute(res, nil)
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("images"))))
	http.HandleFunc("/page/", foo)
	http.ListenAndServe(":8080", nil)
}
