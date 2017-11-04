package main

import(
	_"github.com/go-sql-driver/mysql"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"strings"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"bytes"
	"database/sql"
	"github.com/tidwall/gjson"
)

type Qus struct {
	Question string
	Answer []string
}

var(
	db *sql.DB
	err error
	questions=make([]Qus,0)
	question Qus
)

func init (){
	Json, err := ioutil.ReadFile("config_example.json")
	if err != nil {
		log.Fatal(err)
	}
	JsonStr := string(Json)
	password := gjson.Get(JsonStr, "password")
	port := gjson.Get(JsonStr, "port")
	url := gjson.Get(JsonStr, "url")
	user := gjson.Get(JsonStr, "user")
	pd := password.String()
	pt := port.String()
	ul := url.String()
	ur := user.String()
	db, err = sql.Open("mysql", ur+":"+pd+"@tcp("+ul+":"+pt+")/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
}

func Decode1(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
	if err == nil {
		dst = string(data)
	}
	return
}



func main(){
	doc,err:=goquery.NewDocument("http://www.apesk.com/16pf/")

	if err!= nil {
		log.Fatal(err.Error())
	}

	qs:=make([]string,0)
	as:=make([]string,0)
	doc.Find("td .style1 strong").Each(func(i int,s *goquery.Selection){
		qs=append(qs,Decode1(strings.TrimSpace(s.Text())))
	})
	doc.Find("td .style1 .green").Each(func(i int, s *goquery.Selection) {
		as=append(as,Decode1(strings.TrimSpace(s.Text())))
	})
	for i,q:=range qs {
		question.Question=q
		//与question:=Qus{Question:q}等效
		question.Answer=make([]string,0)
		for j:=i*3;j<i*3+3 ;j++  {
			question.Answer=append(question.Answer, as[j])
		}
		stmt, err := db.Prepare(`INSERT INTO hahaha(question,answer) VALUES(?,?)`)
		//defer stmt.Close()
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec(question.Question,strings.Join(question.Answer,";"))
		if err!=nil {
			log.Fatal(err.Error())
		}
		questions=append(questions,question)
	}

	fmt.Printf("%v\n",questions)
}


