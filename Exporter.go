package main

import (
	"fmt"
	"encoding/csv"
	"encoding/json"
	"encofing/xml"
	"log"
	"os"
	"time"
	"github.com/mmcdole/gofeed"
)

const (
	EXPORTER_VERSION="0.0.2"
	EXPORT_FILE_PREFIX="RssFeedExport-"
)

type rssResource struct {
	rssOrigin string	`json:"origin"`
	rssCategory string	`json:"category"`
	rssURLs []string `json:"urls"` 
}

type rssData struct {
	title string
	link string
	author string
}


func ReadRSS(r *rssResource) ([]rssData) {
	var rssData []rssData
	for _,v := range r.rssURLs {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(v)
		for _, item := range feed.Items {
			line := fmt.Sprintf("%s", item.Title)
			data = append(data,line)
			line = fmt.Sprintf("%s", item.Link)
			data = append(data,line)
			line = fmt.Sprintf("%s", item.Author.Name)
			data = append(data,line)

