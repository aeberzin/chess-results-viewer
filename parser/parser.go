package parser

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	StartList   = "0"
	Pairs       = "2"
	Competitors = "1"
	Table       = "3"
)

func GetTable(tournament string, class string, round ...string) [][]string {
	request := "http://chess-results.com/tnr" + tournament + ".aspx?lan=11&art=" + class
	if len(round) > 0 {
		request += "&rd=" + round[0]
	}
	res, err := http.Get(request)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	table := doc.Find(".CRs1")
	result := [][]string{}
	table.Find("tr").Each(func(_ int, selection *goquery.Selection) {
		row := []string{}
		selection.Find("td").Each(func(_ int, selectionCell *goquery.Selection) {
			row = append(row, selectionCell.Text())
		})
		result = append(result, row)
	})
	return result
}
