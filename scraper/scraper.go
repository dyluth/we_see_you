// find_links_in_page.go
package scraper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ElementProcesser interface {
	ProcessElement(index int, element *goquery.Selection)
}

func scrape(ep ElementProcesser) error {
	// Make HTTP request
	response, err := http.Get("https://www.theyworkforyou.com/mp/24901/steve_brine/winchester/votes")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
		return err
	}

	// /html/body/div[3]/div/div[2]/div[2]/div[1]/ul
	// get all the sections ul class="vote-descriptions"
	sel := document.Find(`ul`) //class="vote-descriptions"
	fmt.Printf("found %v elements\n", sel.Length())
	sel.Each(ep.ProcessElement)
	// inside each should be an <li>, with useful text in the body

	return nil
}

// ScrapeVotingRecord looks at the specific person, and pulls out their voting summaries for each topic
func ScrapeVotingRecord(name string) (map[string]string, error) {

	ep := ElementProcesserImpl{
		results: make(map[string]string),
	}
	scrape(&ep)

	return nil, nil

}

type ElementProcesserImpl struct {
	results map[string]string
}

// This will get called for each HTML element found
func (ep *ElementProcesserImpl) ProcessElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	isVD := element.HasClass("vote-descriptions")
	fmt.Printf("found %v: good? %v \n", index, isVD) //, element.Text())
	if isVD {
		element.Children().Each(ep.ProcessChildElement)
	}
}

func (ep *ElementProcesserImpl) ProcessChildElement(index int, element *goquery.Selection) {
	fmt.Printf("VoteDesc: %v, %v\n", index, element.Text())

	// TODO get the key and the result and put it into ep.results
}
