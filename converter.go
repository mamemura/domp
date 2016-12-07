package main
import (
	"strings"
	"encoding/csv"
	"encoding/json"
	"bytes"
)

type Converter func([][]string) string

func GetConverter(strategy string) Converter{
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
