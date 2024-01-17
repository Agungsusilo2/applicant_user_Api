package main

import (
	"Applicant_user/database"
	"Applicant_user/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.ConnectDatabase()

	apiKey := "Pukulele"
	routes.Api(r, db, apiKey)
	r.Run(":9000")
}
