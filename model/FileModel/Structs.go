package FileModel

import (
	"GoPanClient/model"
)

type File struct {
	ID int `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Url string `gorm:"column:url" json:"url"`
	Creator int `gorm:"column:creator" json:"creator"`
	Modifier int `gorm:"column:modifier" json:"modifier"`
	CreateTime model.BetterTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime model.BetterTime `gorm:"column:update_time" json:"update_time"`
	Type string `gorm:"column:type" json:"type"`
}