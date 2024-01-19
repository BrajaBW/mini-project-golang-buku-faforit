package main

import (
	"net/http"
	"mini-project-golang-buku-faforit/controllers/bukucontroller"
)

func main() {
	http.HandleFunc("/", bukucontroller.Index)
	http.HandleFunc("/buku", bukucontroller.Add)
	http.HandleFunc("/buku/edit", bukucontroller.Edit)
	http.HandleFunc("/buku/delete", bukucontroller.Delete)

	http.ListenAndServe(":4000", nil)
}
