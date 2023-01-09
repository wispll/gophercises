package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	href string
	text string
}

var (
	flagFliePath string
)

func init() {
	flag.StringVar(&flagFliePath, "f", "ex1.html", "html file")
	flag.Parse()
}

func main() {
	f, err := os.Open(flagFliePath)
	if err != nil {
		log.Fatal(err)
	}
    defer f.Close()

	linkSlice := make([]Link, 0, 3)

	z := html.NewTokenizer(f)

	depth := 0
	link := Link{}

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
            fmt.Println("end of file")
            fmt.Printf("%+v\n", linkSlice)
            os.Exit(0)

		case html.TextToken:
			if depth > 0 {
				link.text += string(z.Text())
			}

		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
            /*fmt.Printf("tag name : %s", tn)*/
			if "a" == string(tn) {
				if tt == html.StartTagToken {
					depth++
					if depth == 1 {
						//TODO get attr value
						for {

							key, value, hasNext := z.TagAttr()
							if "href" == string(key) {
								link.href = string(value)
								break
							}
							if !hasNext {
								break
							}
						}
					}
				} else {
					depth--
					if depth == 0 {
						linkSlice = append(linkSlice, link)
                        link.text = ""
                        link.href = ""
					}
				}
			}
		}
	}

}
