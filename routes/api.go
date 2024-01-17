package routes

import (
	applicant "Applicant_user/app/controllers"
	"Applicant_user/middleware"
	"Applicant_user/model/applicant/repository"
	"Applicant_user/model/applicant/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ctx *gin.Context

func Api(router *gin.Engine, db *gorm.DB, apiKey string) {
	applicantRepository := repository.NewApplicantRepository(db)
	applicantService := service.NewApplicantService(applicantRepository)
	applicantController := applicant.NewApplicantController(applicantService, ctx)

	v1 := router.Group("/api/v1", middleware.AuthMiddleware(apiKey))
	{
		v1.GET("/applicantusers", applicantController.Index)
		v1.GET("/applicantusers/:id", applicantController.GetByID)
		v1.POST("/applicantusers", applicantController.Create)
		v1.PATCH("/applicantusers/:id", applicantController.Update)
		v1.DELETE("/applicantusers/:id", applicantController.Delete)
	}
}
