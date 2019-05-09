package spider

import (
		"github.com/PuerkitoBio/goquery"
	"log"
		"regexp"
		"strconv"
		"github.com/guoruibiao/kindlefeeder/errors"
	)

const (
	CSDN_HOST = "http://blog.csdn.net/"
)

type Spider struct{
}

func New() (spider *Spider, err error) {
	return &Spider{}, nil
}

func (spider *Spider) fetch(url string) (*goquery.Document, error) {
	var doc goquery.Document
	document, err := goquery.NewDocument(url)
	if err != nil {
		return &doc, err
	}
	return document, nil
}


// FetchCSDN 抓取文章列表，
// TODO 去重
func (spider *Spider) FetchCSDNPage(url string) ([]string, error) {
	doc, err := spider.fetch(url)
	if err != nil {
		return nil, err
	}
	var links []string
	doc.Find("div").Find(".article-item-box").Each(func(index int, s *goquery.Selection) {
		uri, _ := s.Find("h4").Find("a").Attr("href")
		links = append(links, uri)
	})
	return links, nil
}

func (spider *Spider)FetchURLs(urls[] string) ([]string, error) {
	if spider == nil {
		return nil, &errors.CustomizeError{Name:"spider未初始化"}
	}
	return urls, nil
}

func (spider *Spider)FetchCSDNAuthor(username string)([]string, error) {
	homepage := CSDN_HOST + username
	doc, err := spider.fetch(homepage)
	if err != nil {
		log.Fatal(err)
	}
	// HTML端是JS生成的，所以使用正则抓一下，做个总页数计算
	html, _ := doc.Html()
	pattern, _ := regexp.Compile(`var listTotal = (\d+) ;`)
	submatch := pattern.FindStringSubmatch(html)
	if submatch == nil || len(submatch)<2 {
		return nil, errors.CustomizeError{Name:"没有可下载的博客"}
	}
	totalArticles := submatch[1]
	articles, ok  := strconv.Atoi(totalArticles)
	if ok != nil {
		return nil, err
	}
	pages := 0
	if articles % 20 == 0 {
		pages = articles / 20
	}else{
		pages = articles/20+ 1
	}
	//fmt.Println("total pages: ", pages)

	var alllinks []string
	var pageurl string
	for index:=1; index<=pages; index++ {
		pageurl = CSDN_HOST + username + "/article/list/" + strconv.Itoa(index)
		pagelinks, _ := spider.FetchCSDNPage(pageurl)
		for _, link := range pagelinks {
			alllinks = append(alllinks, link)
		}
	}
    return alllinks, nil
}