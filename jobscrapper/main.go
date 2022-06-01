package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		extractedjobs := getPage(i)
		jobs = append(jobs, extractedjobs...)
	}
	for _, pt_jobs := range jobs {
		fmt.Println(pt_jobs)
	}
	//fmt.Println(jobs)
}

func getPage(page int) []extractedJob {
	var jobs_array []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".cardOutline")

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs_array = append(jobs_array, job)
	})
	return jobs_array
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Find(".jobTitle>a").Attr("data-jk")
	title, _ := card.Find(".jobTitle>a>span").Attr("title")
	location := cleanString(card.Find(".companyLocation").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		summary:  summary,
	}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
