package main

import (
	"fmt"
	"log"

	"github.com/hapi666/playground/crawler"
)

func main() {
	resu, err := crawler.Crawl("http://www.apesk.com/16pf/")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%v\n", resu)
	}
}
