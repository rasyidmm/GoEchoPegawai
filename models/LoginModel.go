package models

import (
	"database/sql"
	"fmt"
	"github.com/rasyidmm/EchoRest/db"
	"github.com/rasyidmm/EchoRest/helpers"
)

type User struct {
	Id int8 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckLogin(username,password string)(bool,error)   {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = $1"

	err := con.QueryRow(sqlStatement,username).Scan(
		&obj.Id,&obj.Username,&pwd)

	if err == sql.ErrNoRows{
		fmt.Println("user tidak ditemukan")
		return false,err
	}
	if err != nil{
		fmt.Println("Query Error")
		return false, err
	}

	match,err := helpers.CheckPasswordHash(password,pwd)
	if!match{
		fmt.Println("Hast and password doesnt match")
		return false,err
	}
	return true,nil
}
