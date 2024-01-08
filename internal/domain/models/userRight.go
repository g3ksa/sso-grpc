package models

type Right struct {
	Code      string     `json:"code"       gorm:"primaryKey;not null;type:varchar(255)"`
	Name      string     `json:"name"       gorm:"type:varchar(255)"`
	GroupCode string     `json:"group_code"`
	Group     RightGroup `json:"group"      gorm:"foreignKey:GroupCode"`
}

func (r *Right) ToShortRight() ShortRight {
	return ShortRight{
		Code:      r.Code,
		GroupCode: r.GroupCode,
	}
}

type ShortRight struct {
	Code      string `json:"code"`
	GroupCode string `json:"group_code"`
}
