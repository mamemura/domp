package main
import (
	"strings"
	"regexp"
)
type Query struct{
	Command string
	Arg string
}

func QueryStrParse(queryStr string) []*Query {
	queryStr = strings.TrimSpace(queryStr)
	tokens := regexp.MustCompile(`\s*\|\s*`).Split(queryStr, -1)
	var queries []*Query
	for _, token := range tokens{
		queries= append(queries, TokenParse(token))
	}
	return queries
}

func TokenParse(token string) *Query {
	token = strings.TrimSpace(token)
	attrs := regexp.MustCompile(`\s*@\s*`).Split(token, 2)
	if len(attrs) == 2{
		return &Query{ Command: attrs[0], Arg: attrs[1] }
	}
	return &Query{ Command: token}
}
