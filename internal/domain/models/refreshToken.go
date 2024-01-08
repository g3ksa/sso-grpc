package models

import "time"

type RefreshToken struct {
	UserId    int       `json:"userId"    gorm:"column:userId;type:int"`
	User      User      `json:"user"      gorm:"foreignKey:UserId"`
	HashedRt  string    `json:"hashedRt"  gorm:"primaryKey;column:hashedRt;type:varchar(255)"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;column:createdAt;type:timestamp"`
}
