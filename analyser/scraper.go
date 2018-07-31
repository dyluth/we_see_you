// find_links_in_page.go
package analyser

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	sections = []string{}
)

func scrape() error {
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
	sel.Each(processElement)
	// inside each should be an <li>, with useful text in the body

	for _, t := range sections {
		fmt.Printf(" sections[\"%v\"]=[]string{}\n", t)
	}

	fmt.Printf(" === TOTAL: %v ===\n", len(sections))
	return nil
}

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	isVD := element.HasClass("vote-descriptions")
	fmt.Printf("found %v: good? %v - %v\n", index, isVD, element.Text())
	if isVD {
		element.Children().Each(processVoteDesc)
	}
}

func processVoteDesc(index int, element *goquery.Selection) {
	fmt.Printf("VoteDesc: %v, %v\n", index, element.Text())
	element.Children().Each(printout)
}

func printout(index int, element *goquery.Selection) {
	text := strings.TrimSpace(element.Text())
	fmt.Printf("Printout: %v, %v\n", index, text)
	if index == 0 {
		sections = append(sections, text)
	} else if index == 1 {
		if !strings.Contains(text, "Show votes") {
			sections[len(sections)-1] = fmt.Sprintf("%v %v", sections[len(sections)-1], text)
		}
	}
}
