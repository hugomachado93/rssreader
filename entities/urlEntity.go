package entities

import "github.com/jinzhu/gorm"

type UrlEntity struct {
	gorm.Model
	Url string `gorm:"unique;not null;"`
}

func NewUrl(url string) *UrlEntity {
	return &UrlEntity{
		Url: url,
	}
}