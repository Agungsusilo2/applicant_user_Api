package applicant

import (
	"Applicant_user/model/applicant/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApplicantController struct {
	applicantService service.ApplicatService
	ctx              *gin.Context
}

func NewApplicantController(applicantService service.ApplicatService, ctx *gin.Context) ApplicantController {
	return ApplicantController{applicantService, ctx}
}

func (a *ApplicantController) Index(ctx *gin.Context) {
	data := a.applicantService.GetAll()

	ctx.JSON(http.StatusOK, data)
}

func (a *ApplicantController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	data := a.applicantService.GetByID(id)

	ctx.JSON(http.StatusOK, data)
}

func (a *ApplicantController) Create(ctx *gin.Context) {
	data, err := a.applicantService.Create(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
			"code":   http.StatusInternalServerError,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (a *ApplicantController) Delete(ctx *gin.Context) {
	data, err := a.applicantService.Delete(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}

func (a *ApplicantController) Update(ctx *gin.Context) {
	data, err := a.applicantService.Update(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"data":   err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, data)
}
