package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"io"
	"github.com/codegangsta/cli"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/crypto/ssh/terminal"
)

func CmdDomCutter(c *cli.Context) error {
	// 初期化
	args := c.Args()
	selector := args.Get(0)
	queryStr := c.String("query")
	outputStr := c.String("output")

	queries := QueryStrParse(queryStr)

	reader := Input()

	doc := MakeDocument(reader)

	res := Analyze(selector, doc, queries)
	conv := GetConverter(outputStr)

	Output(res, conv)

	return nil
}

func MakeDocument(reader io.Reader) *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(reader)
	ErrorLog(err)
	return doc
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

func Analyze(selector string, doc *goquery.Document, queries []*Query) [][]string {
	results := make([][]string, 0)
	doc.Find(selector).Each(func (idx int, sl *goquery.Selection) {
		el := make([]string, 0)
		for _, query := range queries {
			el = append(el, Method(query, sl))
		}
		results = append(results, el)
	})
	return results
}

func ErrorLog( err error){
	if err != nil {
		log.Fatal(err)
	}
}
