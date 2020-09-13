package models

import (
	"github.com/rasyidmm/EchoRest/db"
	"net/http"
)

type Pegawai struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Alamat string `json:"alamat"`
	NoHp string `json:"no_hp"`
}

func FetchAllPegawai()(Response,error)  {
 var obj Pegawai
 var arrobj []Pegawai
 var res Response
 con := db.CreateCon()
 sqlStatement := "select * from pegawai"

 rows,err := con.Query(sqlStatement)
 defer rows.Close()
 if err != nil{
 	return res, err
 }

 for rows.Next(){
 	err = rows.Scan(&obj.Id, &obj.Name,&obj.Alamat,&obj.NoHp)
	 if err != nil{
		 return res, err
	 }

	 arrobj= append(arrobj,obj)
 }
 res.Status = http.StatusOK
 res.Message="Sukses"
 res.Data=arrobj

 return res,nil
}

func StorePegawai(id int,nama string,alamat string,nohp string)(Response,error)  {
	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT INTO pegawai (id,nama, alamat, no_hp)VALUES ($1, $2, $3, $4)"
	stmt,err := con.Prepare(sqlStatement)

	if err != nil{
		return res,err
	}
	result, err := stmt.Exec(id,nama,alamat,nohp)
	if err != nil{
		return  res,err
	}
	lastInsertId,err := result.LastInsertId()
	if err != nil{
		return  res,err
	}
	res.Status = http.StatusOK
	res.Message="Sukses"
	res.Data= map[string]int64{
		"last_inserted_id": lastInsertId,
	}

	return  res,nil
}
func UpdatePegawai(id int,nama string,alamat string,nohp string) (Response,error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "UPDATE pegawai SET nama=$1, alamat=$2,	 no_hp=$3 where id=$4"
	stmt,err := con.Prepare(sqlStatement)
	if err != nil{
		return res,err
	}
	result, err := stmt.Exec(nama,alamat,nohp,id)
	if err != nil{
		return  res,err
	}
	rowsAffected,err := result.RowsAffected()
	if err != nil{
		return  res,err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affectes": rowsAffected,
	}
	return res, nil
}

func DeletePegawai(id int)(Response, error)  {

	var res Response
	con := db.CreateCon()
	sqlStatement := "DELETE FROM pegawai where id=$1"
	stmt,err := con.Prepare(sqlStatement)
	if err != nil{
		return res,err
	}
	result, err := stmt.Exec(id)
	if err != nil{
		return  res,err
	}
	rowsAffected,err := result.RowsAffected()
	if err != nil{
		return  res,err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affectes": rowsAffected,
	}
	return res, nil
}