package controller

import (
	"rssreader/service"
)

type RssReaderController struct {
	service.RssReaderService
}

func (rssReaderController RssReaderController) commandos() {
}