package main

import (
	"hacktiv8-chapter-two-project-one/database"
	"hacktiv8-chapter-two-project-one/routes"

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
