package main

import (
	"fmt"
	"log"

	"github.com/hapi666/Go-Crawler/crawler"
)

func main() {
	resu, err := crawler.Crawler("http://www.apesk.com/16pf/")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%v\n", resu)
}
