package main

import "fmt"
import "log"
import "os"
import "encoding/csv"
import "github.com/mmcdole/gofeed"

const ExportFile = "FeedExport.csv"

var rss map[string]string

func main() {

	/* greeting :P */
	log.Println("Welcome to V1k3r14n feeds reader. This stuff is based mainly on https://github.com/mmcdole/gofeed module. Enjoy...")
	log.Println("Initializing feed databank...")

	/* which feeds to read*/
	rss := map[string]string{
		"articles":          "https://www.cshub.com/rss/articles",
		"demos":             "https://www.cshub.com/rss/demos",
		"news":              "https://www.cshub.com/rss/news",
		"whitepapers":       "https://www.cshub.com/rss/whitepapers",
		"attacks":           "https://www.cshub.com/rss/categories/attacks",
		"case-studies":      "https://www.cshub.com/rss/categories/case-studies",
		"iot":               "https://www.cshub.com/rss/categories/iot",
		"malware":           "https://www.cshub.com/rss/categories/malware",
		"mobile":            "https://www.cshub.com/rss/categories/mobile",
		"network":           "https://www.cshub.com/rss/categories/network",
		"security-strategy": "https://www.cshub.com/rss/categories/security-strategy",
		"threat-defense":    "https://www.cshub.com/rss/categories/threat-defense",
	}

	/* open export file */
	log.Println("Opening export file...")
	file, err := os.Create(ExportFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	/* Create file writer instance */
	writer := csv.NewWriter(file,)
	defer writer.Flush() /* this maybe could be inside rss loop, not sure, but it should be more memory friendly */

	/* iterate over feed, create data 2 dim slices and write them to file */
	log.Println("Now we are reading freeds and export them to file " + ExportFile)
	var data []string
	data = append(data,"Title")
	data = append(data,"Link")
	data = append(data,"Author")
	writer.Write(data)
	data = data[:0]
	for _, v := range rss {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(v)
		for _, item := range feed.Items {
			line := fmt.Sprintf("%s,", item.Title)
			data = append(data,line)

			line = fmt.Sprintf("%s,", item.Link)
			data = append(data,line)

			line = fmt.Sprintf("%s,", item.Author.Name)
			data = append(data, line)
			writer.Write(data)
			data = data[:0]
		}
	}
	log.Println("Everything should be writed into .csv. Enjoy...")
}
