package models

type Role struct {
	Code  string `json:"code" gorm:"primaryKey;unique;not null"`
	Name  string `json:"name"`
	Users []User `json:"users" gorm:"foreignKey:RoleCode"`
}
