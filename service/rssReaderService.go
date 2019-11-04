package service

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"rssreader/model"
	"rssreader/repository"
)

type RssReaderService struct {
	repository.RssRepository
}

func (rssReader * RssReaderService) ReadRss(url string) {
	//fp := gofeed.NewParser()
	//feed, _ := fp.ParseURL(url)
}

func (rssReader *RssReaderService) SaveUrl(url string){
	err := rssReader.RssRepository.Save(url)
	if err != nil {
		fmt.Println(err)
	}
}

func (rssReader *RssReaderService) GetUrls() []model.UrlModel {
	result, err := rssReader.RssRepository.FindAll()

	if err != nil {
		panic(err)
	}

	for _, element := range result{
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(element.Url)

		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(element.Index, element.Url)
		fmt.Println()
		fmt.Println(feed.Title)
		fmt.Println(feed.Author)
		fmt.Println(feed.Categories)
		fmt.Println(feed.Published)
	}

	return result
}

func (rssReader * RssReaderService) ShowArticle(id uint, showArticle bool, showAll bool) {

	url := rssReader.RssRepository.FindById(id)

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)

	if err != nil {
		fmt.Println(err)
	}

	if showAll {
		feed.String()
	}
	if showArticle {
		fmt.Println(feed.Description)
	} else {

		fmt.Println(feed.Title)
		fmt.Println(feed.Author)
		fmt.Println(feed.Categories)
		fmt.Println(feed.Published)
	}

}