package main

import (
	"net/http"
	"mini-project-golang-buku-faforit/controllers/homecontroler"
	"mini-project-golang-buku-faforit/controllers/bukucontroler"
)

func main() {
	// 1.homepage
	http.HandleFunc("/", homecontroler.Index)
	// 2.buku
	http.HandleFunc("/buku/index", bukucontroler.Index)
	http.HandleFunc("/buku",bukucontroler.Add)
	http.HandleFunc("/buku/edit", bukucontroler.Edit)
	http.HandleFunc("/buku/delete", bukucontroler.Delete)

	http.ListenAndServe(":4000", nil)
}
