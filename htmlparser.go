package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"io"
	"regexp"
	"github.com/codegangsta/cli"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/crypto/ssh/terminal"
)

func CmdHtmlParser(c *cli.Context) error {
	args := c.Args()
	selector := args.Get(0)
	queryStr := c.String("query")


	queries := QueryParse(queryStr)

	reader := Input()
	doc := Parse(reader)
	res := Analyze(selector, doc, queries)
	Output(res)
	return nil
}

func Parse(reader io.Reader) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(reader)
	ErrorLog(err)
	return doc
}

func QueryParse(queryStr string) []string {
	queryStr = strings.TrimSpace(queryStr)
	queries := regexp.MustCompile(`\s*\|\s*`).Split(queryStr, -1)
	return queries
}

func Input() io.Reader{
	if terminal.IsTerminal(syscall.Stdin) {
		var stdin string
		fmt.Scan(&stdin)
		return strings.NewReader(stdin)
	}
	return os.Stdin
}

func Output(data [][]string) {
	for _, values := range data {
		fmt.Println(strings.Join(values, ","))
	}
}


func Analyze(selector string, doc *goquery.Document, queries []string) [][]string {
	results := make([][]string, 0)
	doc.Find(selector).Each(func (idx int, sl *goquery.Selection) {
		el := make([]string, 0)
		for _, query := range queries {
			switch {
			case query == "text":
				el = append(el, Text(sl))
			case regexp.MustCompile("attr@.*").MatchString(query):
				attrs := regexp.MustCompile(`\s*@\s*`).Split(query, 2)
				el = append(el, Attr(attrs[1], sl))
			case query == "html": fallthrough;
			default:
				el = append(el, Html(sl))
			}
		}
		results = append(results, el)
	})
	return results
}

func Html(selector *goquery.Selection) string{
	html, err := selector.Html()
	ErrorLog(err)
	return strings.TrimSpace(html)
}

func Text(selector *goquery.Selection) string{
	text := selector.Text()
	return strings.TrimSpace(text)
}

func Attr(attr string, selector *goquery.Selection) string {
	value, exists := selector.Attr(attr)
	if !exists {
		value = ""
	}
	return strings.TrimSpace(value)
}

func ErrorLog( err error){
	if err != nil {
		log.Fatal(err)
	}
}
