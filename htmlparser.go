package main

// TODO
// * メソッドを機能ごとに分ける。
// * 自分の構造とかもみたいよね

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
	"encoding/csv"
	"encoding/json"
	"bytes"
)

func CmdHtmlParser(c *cli.Context) error {
	args := c.Args()
	selector := args.Get(0)
	queryStr := c.String("query")
	outputStr := c.String("output")

	queries := QueryParse(queryStr)

	reader := Input()
	doc := Parse(reader)
	res := Analyze(selector, doc, queries)
	strtagey := selectStrategy(outputStr)
	Output(res, strtagey)
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

func Output(data [][]string, conv Converter) {
	fmt.Printf(conv(data))
}

type Converter func([][]string) string

func TextConverter(data [][]string) string {
	strs := make([]string, 0)
	for _, values := range data {
		strs = append(strs,strings.Join(values, " "))
	}
	return strings.Join(strs, "\n")
}

func CsvConverter(data [][]string) string {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	w.WriteAll(data)
	return buf.String()
}

func JsonConverter(data [][]string) string {
	buf, _ :=json.Marshal(data)
	return string(buf)
}

func selectStrategy(strategy string) Converter{
	switch strategy {
	case "csv":
		return CsvConverter
	case "json":
		return JsonConverter
	case "text": fallthrough;
	default:
		return TextConverter
	}
}

func Analyze(selector string, doc *goquery.Document, queries []string) [][]string {
	results := make([][]string, 0)
	doc.Find(selector).Each(func (idx int, sl *goquery.Selection) {
		el := make([]string, 0)
		for _, query := range queries {
			switch {
			case query == "html":
				el = append(el, Html(sl))
			case regexp.MustCompile("attr@.*").MatchString(query):
				attrs := regexp.MustCompile(`\s*@\s*`).Split(query, 2)
				el = append(el, Attr(attrs[1], sl))
			case query == "text": fallthrough;
			default:
				el = append(el, Text(sl))
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
