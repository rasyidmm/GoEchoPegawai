package routes

import (
	"github.com/labstack/echo"
	"github.com/rasyidmm/EchoRest/controllers"
	"github.com/rasyidmm/EchoRest/middleware"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK,"Hello this is Echo")
	})

	e.GET("/pegawai",controllers.FetchAllPegawai, middleware.IsAuthenticated)

	e.POST("/pegawai",controllers.StorePegawai)

	e.PUT("/pegawai",controllers.UpdatePegawai)

	e.DELETE("/pegawai",controllers.DeletePegawai)

	e.GET("/generate-hash/:password",controllers.GenerateHashPassword)
	e.POST("/login",controllers.CheckLogin)
	return e
}
