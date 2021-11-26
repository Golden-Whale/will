package model

import (
	"gorm.io/gorm"
	"will/util"
)

type Will struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt util.TimeStamp
	UpdatedAt util.TimeStamp `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
