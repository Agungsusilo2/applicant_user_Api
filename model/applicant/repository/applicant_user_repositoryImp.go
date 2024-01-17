package repository

import (
	"Applicant_user/model/applicant/domain"
	"gorm.io/gorm"
)

type ApplicantRepositoryImpl struct {
	db *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) ApplicantRepository {
	return &ApplicantRepositoryImpl{db}
}
func (ar *ApplicantRepositoryImpl) FindAll() []domain.ApplicantUser {
	var applicants []domain.ApplicantUser

	_ = ar.db.Find(&applicants)

	return applicants
}

func (ar *ApplicantRepositoryImpl) FindById(id string) domain.ApplicantUser {
	var applicant domain.ApplicantUser

	_ = ar.db.First(&applicant, "id = ?", id)

	return applicant
}

func (ar *ApplicantRepositoryImpl) Save(applicant domain.ApplicantUser) (*domain.ApplicantUser, error) {
	result := ar.db.Create(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}

func (ar *ApplicantRepositoryImpl) Update(applicant domain.ApplicantUser) (*domain.ApplicantUser, error) {
	result := ar.db.Save(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}

func (ar *ApplicantRepositoryImpl) Delete(applicant domain.ApplicantUser) (*domain.ApplicantUser, error) {
	result := ar.db.Delete(&applicant)

	if result.Error != nil {
		return nil, result.Error
	}

	return &applicant, nil
}
