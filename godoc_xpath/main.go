package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/xquery/html"
)

type result struct {
	url        string
	importPath string
}

type results []*result

func (rs results) ToString() string {
	var representatives []string
	for _, r := range rs {
		representatives = append(
			representatives,
			fmt.Sprintf(
				"(url: %s, importPath: %s)",
				r.url,
				r.importPath,
			),
		)
	}
	return strings.Join(representatives, ",")
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s query", os.Args[0])
	}

	url := fmt.Sprintf("https://godoc.org/?q=%s", os.Args[1])
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to get result: %s", err)
	}
	defer res.Body.Close()

	doc, err := htmlquery.Parse(res.Body)
	if err != nil {
		log.Fatalf("failed to parse response body: %s", err)
	}

	var rs results
	for _, n := range htmlquery.Find(doc, "/html/body/div/table/tbody/tr/td/a") {
		rs = append(
			rs,
			&result{
				url:        htmlquery.InnerText(n),
				importPath: htmlquery.SelectAttr(n, "href"),
			},
		)
	}

	fmt.Printf(rs.ToString())
}
