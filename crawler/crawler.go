package crawler

import (
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

func ProcessDate(doc *goquery.Document) {

}

func Crawler(url string) ([]Qus, error) {
	var (
		questions = make([]Qus, 0)
		qs        string
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	r := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	doc.Find("form .style1 .style1").Each(func(i int, s *goquery.Selection) {
		ans := make([]string, 0)
		s.Find("strong").Each(func(i int, selection *goquery.Selection) {
			qs = strings.TrimSpace(s.Text())
		})
		s.Find(".green").Each(func(i int, selection *goquery.Selection) {
			ans = append(ans, strings.TrimSpace(s.Text()))
		})
		questions = append(questions, Qus{qs, ans})

	})
	return questions, nil
}
