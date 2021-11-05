package service

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const _rssUrl = "https://habrahabr.ru/rss/best/"

type habrRss struct {
	Items []habrItem `xml:"channel>item"`
}

type habrItem struct {
	Title string `xml:"title"`
}

type HabrDownloaderService struct {
}

func ProvideHabrDownloaderService() *HabrDownloaderService {
	return &HabrDownloaderService{}
}

func (service *HabrDownloaderService) DownloadBestTitles() ([]string, error) {
	response, err := http.Get(_rssUrl)
	if err != nil {
		return nil, errors.Wrap(err, "habr downloader service")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "habr downloader service")
	}

	rss := &habrRss{}
	if err = xml.Unmarshal(body, rss); err != nil {
		return nil, errors.Wrap(err, "habr downloader service")
	}

	titles := make([]string, 0, len(rss.Items))
	for _, item := range rss.Items {
		titles = append(titles, item.Title)
	}
	return titles, nil
}
