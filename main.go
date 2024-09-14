package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main()  {
	blogTitles, err:= getLatestBlogTitles("https://golangcode.com")   //calling our blog title function (but we haven't written it yet)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("Blong Titles: ")
	fmt.Print(blogTitles)
	
}

func getLatestBlogTitles(url string)(string, error)  {
	resp, err:=http.Get(url)
	if err != nil {
		return "", err
	}
	doc, err:= goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	titles:= ""
	doc.Find(".post-title").Each(func (i int, s *goquery.Selection)  {
		titles += "-" + s.Text() + "\n" 
	})
		return titles, nil
}