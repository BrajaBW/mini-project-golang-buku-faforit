package bukucontroler

import (
	"html/template"
	"mini-project-golang-buku-faforit/model/bukumodel"
	"net/http"
)
func Index(respon http.ResponseWriter, request *http.Request) {
	bukues := bukumodel.GetAll()
	data := map[string]any{
		"bukus":bukues,
	}

	temp, err := template.ParseFiles("views/buku/index.html")
	if err != nil{
		panic(err)
	}
		
	temp.Execute(respon, data)
	


}
func Add(respon http.ResponseWriter, request *http.Request) {

}

func Edit(respon http.ResponseWriter, request *http.Request) {

}

func Delete(respon http.ResponseWriter, request *http.Request) {

}