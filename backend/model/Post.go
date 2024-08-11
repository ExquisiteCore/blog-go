package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:TEXT" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
	Cid      int64    `gorm:"type:int;not null" json:"cid"`
}
