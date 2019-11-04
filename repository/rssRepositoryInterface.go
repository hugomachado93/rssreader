package repository

import (
	"rssreader/model"
)

type RssRepository interface {
	FindAll() ([]model.UrlModel, error)
	Save(url string) error
	FindByUrl(urlName string) string
	FindById(id uint) string
}