package domain

type ApplicantUser struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Fullname string `json:"full_name" gorm:"type:varchar(255);not null"`
	Job      string `json:"job" gorm:"type:varchar(255);not null"`
	Gmail    string `json:"gmail" gorm:"type:varchar(255);not null"`
	Telepon  int    `json:"telepon" gorm:"type:int;not null"`
	Age      int    `json:"age" gorm:"type:int;not null"`
	Username string `json:"username" gorm:"type:varchar(255);not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}
