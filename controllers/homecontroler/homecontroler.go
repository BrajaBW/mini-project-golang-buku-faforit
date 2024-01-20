package homecontroler

import (
	"net/http"
	"html/template"
)

func Index(respon http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/buku/index.html")
	if err != nil{
		panic(err)
	}

	temp.Execute(respon, nil)

}

