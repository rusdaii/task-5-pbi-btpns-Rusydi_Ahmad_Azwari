package main

import (
	"fmt"
	"go-project/database"
	"go-project/helpers"
	"go-project/router"
)

func main() {
	helpers.LoadEnvVariables()

	database.InitDB()

	database.Migrate()
	
	r := router.InitRouter()

	r.Run(fmt.Sprintf(":%d", router.Port()))  
}