package controllers

import (
	"github.com/labstack/echo"
	"github.com/rasyidmm/EchoRest/models"
	"math/rand"
	"net/http"
	"strconv"
)

func FetchAllPegawai(c echo.Context) error{
	result, err := models.FetchAllPegawai()
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{"message":err.Error()})
	}

	return c.JSON(http.StatusOK,result)
}
func StorePegawai(c echo.Context)error{
	id := rand.Int()
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	nohp := c.FormValue("nohp")

	result,err :=models.StorePegawai(id,nama,alamat,nohp)

	if err != nil{
		return  c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,result)
}

func UpdatePegawai(c echo.Context)error  {
	id  :=c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	nohp := c.FormValue("nohp")

	convId,err := strconv.Atoi(id)
	if err != nil{
		return  c.JSON(http.StatusInternalServerError,err.Error())
	}
	result,err :=models.UpdatePegawai(convId,nama,alamat,nohp)
	if err != nil{
		return  c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,result)
}

func DeletePegawai(c echo.Context)error  {
	id  :=c.FormValue("id")
	convId,err := strconv.Atoi(id)
	if err != nil{
		return  c.JSON(http.StatusInternalServerError,err.Error())
	}
	result,err :=models.DeletePegawai(convId)
	if err != nil{
		return  c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,result)
}