package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"rssreader/cmd"
	"rssreader/database"
	"rssreader/entities"
	"rssreader/repository"
	"rssreader/service"
)

func main(){

	db := database.ConnectDatabase()
	db.InitializeDatabase(&entities.UrlEntity{})
	defer db.CloseConnection()

	repository := repository.NewRssRepository()
	service := service.RssReaderService{repository}
	cmd.NewCmd(service)

	cmd.Execute()
}