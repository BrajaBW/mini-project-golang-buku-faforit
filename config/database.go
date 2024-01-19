package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func DBConection() (*sql.DB,error){
	dbDriver := "postgres"
	dbUser := "postgres"
	dbPass := "root"
	dbName := "go_buku_fav"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
