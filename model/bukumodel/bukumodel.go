package bukumodel

import (
		"mini-project-golang-buku-faforit/config"
	"mini-project-golang-buku-faforit/entities"
)

func GetAll()[]entities.Buku{
	rows, err := config.DB.Query("SELECT * FROM buku")
	if err != nil{
		panic(err)
	}
	
	defer rows.Close()

	var bukujenis []entities.Buku

	for rows.Next(){
		var buku entities.Buku
		if err := rows.Scan(&buku.Id, &buku.JudulBuku, &buku.Penulis, &buku.Rating);err != nil{
			panic(err)
		}
	
		bukujenis = append(bukujenis, buku)
	}

return bukujenis

}