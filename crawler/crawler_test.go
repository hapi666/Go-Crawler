package crawler

import (
	"testing"
)

func TestCrawler(t *testing.T) {
	test := "http://www.apesk.com/16pf/"
	resultTest, err := Crawler(test)
	if err != nil {
		t.Errorf("crawler(): %v", err)
	} else if len(resultTest) != 187 {
		t.Error("error:Inaccurate data")
	}
}
