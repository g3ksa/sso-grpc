package models

type RightGroup struct {
	Code string `json:"code" gorm:"primaryKey;unique;not null;type:varchar(255)"`
	Name string `json:"name" gorm:"type:varchar(255)"`
}
