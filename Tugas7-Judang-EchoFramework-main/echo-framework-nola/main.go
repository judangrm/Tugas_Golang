package main

import (
	db "echo-framework-juda/db"
	routes "echo-framework-juda/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))
}
