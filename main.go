package main

import (
	"github.com/rasyidmm/EchoRest/db"
	"github.com/rasyidmm/EchoRest/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":1234"))
}
