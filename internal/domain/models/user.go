package models

import (
	"time"
)

type User struct {
	Id         int            `json:"id,omitempty"         gorm:"primaryKey;autoIncrement"`
	CreatedAt  time.Time      `json:"createdAt,omitempty"  gorm:"autoCreateTime;column:createdAt;type:timestamp"`
	ImageUrl   string         `json:"imageUrl,omitempty"   gorm:"column:imageUrl;type:varchar(255)"`
	HashedRts  []RefreshToken `json:"hashedRts,omitempty"  gorm:"foreignKey:UserId`
	Email      string         `json:"email"      gorm:"type:varchar(255)"`
	Password   string         `json:"password"`
	FirstName  string         `json:"firstName"  gorm:"column:firstName;type:varchar(255)"`
	LastName   string         `json:"lastName"   gorm:"column:lastName;type:varchar(255)"`
	IsActive   bool           `json:"isActive,omitempty"   gorm:"column:isActive"`
	Company    string         `json:"company"    gorm:"column:companyCode;type:varchar(255)"`
	RoleCode   string         `json:"roleCode" gorm:"column:roleCode;type:varchar(255)"`
	Role       Role           `json:"role"      gorm:"foreignKey:RoleCode"`
	UserRights []Right        `json:"userRights,omitempty" gorm:"many2many:user_rights;"`
}
