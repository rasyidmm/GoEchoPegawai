package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/rasyidmm/EchoRest/helpers"
	"github.com/rasyidmm/EchoRest/models"
	"net/http"
	"time"
)

func CheckLogin(e echo.Context)error  {
	username :=  e.FormValue("username")
	password := e.FormValue("password")

	res, err := models.CheckLogin(username,password)
	if err != nil{
		return e.JSON(http.StatusInternalServerError,map[string]string{
			"message": err.Error(),
		})
	}
	if !res{
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)


	claims := token.Claims.(jwt.MapClaims)
	claims["username"]=username
	claims["level"]="application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError,map[string]string{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}
func GenerateHashPassword(e echo.Context)error{
	password := e.Param("password")
	hash,_ :=helpers.HashPassword(password)

	return e.JSON(http.StatusOK,hash)
	
}

