package repository

import "Applicant_user/model/applicant/domain"

type ApplicantRepository interface {
	FindAll() []domain.ApplicantUser
	FindById(id string) domain.ApplicantUser
	Save(applicant domain.ApplicantUser) (*domain.ApplicantUser, error)
	Update(applicant domain.ApplicantUser) (*domain.ApplicantUser, error)
	Delete(applicant domain.ApplicantUser) (*domain.ApplicantUser, error)
}
