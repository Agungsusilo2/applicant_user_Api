package service

import (
	"Applicant_user/model/applicant/domain"
	"github.com/gin-gonic/gin"
)

type ApplicatService interface {
	GetAll() []domain.ApplicantUser
	GetByID(id string) domain.ApplicantUser
	Create(ctx *gin.Context) (*domain.ApplicantUser, error)
	Update(ctx *gin.Context) (*domain.ApplicantUser, error)
	Delete(ctx *gin.Context) (*domain.ApplicantUser, error)
}
