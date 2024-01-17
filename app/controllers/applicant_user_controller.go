package applicant

import "github.com/gin-gonic/gin"

type ApplicantUserController interface {
	Index(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}
