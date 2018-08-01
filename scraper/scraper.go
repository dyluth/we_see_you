package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ElementProcesser is the interface for how we process pages.
type ElementProcesser interface {
	ProcessElement(index int, element *goquery.Selection)
}

// Relevent info for a person to look up in theyworkforyou
type Person struct {
	ID           string
	Name         string
	Constituency string
}

func scrape(ep ElementProcesser, person Person) error {
	// Make HTTP request
	urlString := fmt.Sprintf("https://www.theyworkforyou.com/mp/%v/%v/%v/votes", person.ID, person.Name, person.Constituency)
	response, err := http.Get(urlString)
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

// ElementProcesserImpl processes the pages to get the good info out of it.
type ElementProcesserImpl struct {
	results map[string]string
}

// ProcessElement will get called for each `ul` element found
func (ep *ElementProcesserImpl) ProcessElement(index int, element *goquery.Selection) {
	// we are only interested in the vote-descriptions ones
	isVD := element.HasClass("vote-descriptions")
	fmt.Printf("found %v: good? %v \n", index, isVD) //, element.Text())
	if isVD {
		element.Children().Each(ep.ProcessChildElement)
	}
}

// ProcessChildElement is the bit that looks at the voting section and pulls out the useful information
func (ep *ElementProcesserImpl) ProcessChildElement(index int, element *goquery.Selection) {
	fmt.Printf("VoteDesc: %v, %v\n", index, element.Text())
	// get all the text before the `Show votes` bit
	split := strings.Split(element.Text(), "Show votes")
	if len(split) != 2 {
		fmt.Printf("ERROR! splitting text didnt split on `Show votes` into 2 bits but `%v`\n", len(split))
		return
	}
	voteSummary := strings.TrimSpace(split[0])

	// we build sectionname out of the children parts
	sectionName := ""
	element.Children().Each(func(index int, item *goquery.Selection) {
		text := strings.TrimSpace(item.Text())
		if index == 0 {
			sectionName = text
		} else if index == 1 {
			if !strings.Contains(text, "Show votes") {
				// if there are multiple parts before `show votes` append them together
				sectionName = fmt.Sprintf("%v %v", sectionName, text)
			}
		}
	})
	ep.results[sectionName] = voteSummary
}
