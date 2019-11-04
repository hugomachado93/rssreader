package repository

import (
	"errors"
	"fmt"
	"rssreader/database"
	"rssreader/entities"
	"rssreader/model"
)

type RssRepositoyrImpl struct {

}

func NewRssRepository() RssRepository {
	return &RssRepositoyrImpl{}
}

func (rss *RssRepositoyrImpl) FindAll() ([]model.UrlModel, error) {
	urls := []entities.UrlEntity{}

	db := database.ConnectDatabase()
	err := db.GetConnection().Find(&urls).Error

	if err != nil {
		panic(err)
	}

	urlsModel := make([]model.UrlModel, len(urls))

	for index, element := range urls {
		urlsModel[index] = model.UrlModel{Index: element.ID, Url:element.Url}
	}

	return urlsModel, nil

}

func (rss *RssRepositoyrImpl) Save(url string) error{

	if url == "" {
		return errors.New("Url cant be empty")
	}

	urlEntity := entities.NewUrl(url)

	db := database.ConnectDatabase()
	err := db.GetConnection().Save(&urlEntity).Error

	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (rss *RssRepositoyrImpl) FindByUrl(urlName string) string {

	url := entities.UrlEntity{}

	db := database.ConnectDatabase()

	row := db.GetConnection().Model(&url).Where("Url = ?", urlName).Select("Url").Row()

	var result string

	row.Scan(&result)

	return result
}

func (rss *RssRepositoyrImpl) FindById(id uint) string {

	url := entities.UrlEntity{}

	db := database.ConnectDatabase()

	row := db.GetConnection().Model(&url).Where("id = ?", id).Select("Url").Row()

	var result string

	row.Scan(&result)

	return result

}