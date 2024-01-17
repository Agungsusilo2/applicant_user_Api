package dto

type UpdateApplicantReq struct {
	Fullname string `json:"full_name" validate:"required,max=50,min=1"`
	Gmail    string `json:"gmail" validate:"email,required,max=50,min=1"`
	Job      string `json:"job" validate:"required,max=25,min=1"`
	Telepon  int    `json:"telepon" validate:"required,numeric,min=1"`
	Age      int    `json:"age" validate:"required,gte=16,lte=100,numeric"`
	Username string `json:"username" validate:"required,min=4,max=20,alphanum"`
	Password string `json:"password" validate:"required,min=8,max=30,alphanum"`
}
