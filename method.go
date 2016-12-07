package main
import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func Method(query *Query, sl *goquery.Selection) string {
	switch query.Command {
	case "html":
		return Html(sl)

	case "outerhtml":
		return OuterHtml(sl)

	case "nodename":
		return NodeName(sl)

	case "attrs":
		return Attrs(sl)

	case "attr":
		return Attr(query.Arg, sl)

	case "text": fallthrough;
	default:
		return Text(sl)
	}
}


func Html(selector *goquery.Selection) string{
	html, err := selector.Html()
	ErrorLog(err)
	return strings.TrimSpace(html)
}

func OuterHtml(selector *goquery.Selection) string{
	html, err := goquery.OuterHtml(selector)
	ErrorLog(err)
	return strings.TrimSpace(html)
}

func Text(selector *goquery.Selection) string{
	text := selector.Text()
	return strings.TrimSpace(text)
}

func Attrs(selector *goquery.Selection) string {
	node := selector.Get(0)
	var data []string
	for _, attr := range node.Attr {
		data = append(data, attr.Key + "=>" + attr.Val)
	}
	// node.Attr

	return strings.TrimSpace(strings.Join(data, ","))
}

func Attr(attr string, selector *goquery.Selection) string {
	value, exists := selector.Attr(attr)
	if !exists {
		value = ""
	}
	return strings.TrimSpace(value)
}

func NodeName(selector *goquery.Selection) string {
	return goquery.NodeName(selector)
}
