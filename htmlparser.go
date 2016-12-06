package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	_ "bufio"
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
	outputmode := c.String("output")

	reader := Input()
	doc := Parse(reader)
	Output(selector, doc, outputmode)
	return nil
}

func Parse(reader io.Reader) *goquery.Document {
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

func Output(selector string, doc *goquery.Document, mode string)  {
	doc.Find(selector).Each(func (idx int, doc *goquery.Selection) {
		switch {
		case mode == "text":
			Text(doc)
			break
		case regexp.MustCompile("attr@.*").MatchString(mode):
			attr := regexp.MustCompile(`\s*@\s*`).Split(mode, 2)
			Attr(attr[1], doc)
			break
		case mode == "html":
		default:
			Html(doc)
		}
	})
}

func Html(selector *goquery.Selection){
	html, err := selector.Html()
	ErrorLog(err)
	fmt.Println(html)
}

func Text(selector *goquery.Selection){
	text := selector.Text()
	fmt.Println(text)
}

func Attr(attr string, selector *goquery.Selection){
	value, exists := selector.Attr(attr)
	if exists {
		fmt.Println(value)
	}
}

func ErrorLog( err error){
	if err != nil {
		log.Fatal(err)
	}
}
