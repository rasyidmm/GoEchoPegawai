package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rasyidmm/EchoRest/config"
)

var db *sql.DB
var err error

func Init()  {
	conf := config.GetConfig()
	connectionString := "postgres://"+conf.DB_USERNAME+":"+conf.DB_PASSWORD+"@"+conf.DB_HOST+":5431/"+conf.DB_NAME+"?sslmode=disable"
	db ,err = sql.Open("postgres",connectionString)
	fmt.Println(connectionString)
	if err != nil{
		panic("Db Error")
	}

	err = db.Ping()
	if err !=nil{
		panic("DSN invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}