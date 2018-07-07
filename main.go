package main

import (
	"fmt"
	"log"

	"crawler/crawler"
)

func main() {
	reader, err := crawler.Crawler("http://www.apesk.com/16pf/")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		resu := crawler.ProcessDate(reader)
		fmt.Printf("%v\n", resu)
	}
}
