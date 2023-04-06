package main

import (
	"project-2/database"
	"project-2/routes"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	router := gin.Default()

	database.StartDB()
	db := database.GetDB()

	routes.SetupBookRoute(router, db)

	router.Run(PORT)
}
