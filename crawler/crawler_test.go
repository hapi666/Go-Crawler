package crawler

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func mustReadTestData(name string) []byte {
	b, err := ioutil.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		panic(err)
	}
	return b
}

func TestProcessData(t *testing.T) {
	tests := []struct {
		want  []Qus
		input []byte
	}{
		{
			[]Qus{{Question: "1.我很明了本测试的说明：", Answer: []string{"A.是的", "B.不一定", "C.不是的"}}, {Question: "2.我对本测试的每一个问题，都能做到诚实地回答：", Answer: []string{"A.是的", "B.不一定", "C.不是的"}}},
			mustReadTestData("testdata.html"),
		},
	}
	for _, test := range tests {

		testReader := bytes.NewReader(test.input)
		got, err := processData(testReader)
		if err != nil {
			t.Errorf("ProcessDate(%v) return err = %v", testReader, err)
			continue
		}
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("ProcessDate(%v) = %v was incorrect; want %v", test.input, got, test.want)
		}
	}

}

func TestCrawl(t *testing.T) {
	tests := []string{
		"http://www.baidu.com",
		"http://www.github.com",
	}
	for _, test := range tests {
		_, err := Crawl(test)
		if err != nil {
			t.Errorf("Crawler(%s) return error was %v", test, err)
		}
	}
}
