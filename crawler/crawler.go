package crawler

import (
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Qus struct {
	Question string
	Answer   []string
}

func processData(reader io.Reader) ([]Qus, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var (
		questions []Qus
		qs        string
	)
	doc.Find("form .style1 .style1").Each(func(i int, s *goquery.Selection) {
		ans := make([]string, 0)
		s.Find("strong").Each(func(i int, selection *goquery.Selection) {
			qs = strings.TrimSpace(selection.Text())
		})
		s.Find(".green").Each(func(i int, selection *goquery.Selection) {
			ans = append(ans, strings.TrimSpace(selection.Text()))
		})
		questions = append(questions, Qus{qs, ans})
	})
	return questions, nil
}

func Crawl(url string) ([]Qus, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	r := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	result, err := processData(r)
	return result, nil
}
