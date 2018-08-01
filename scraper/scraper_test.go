package scraper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/require"
)

var (
	sections = []string{}
)

func TestScraper(t *testing.T) {

	ep := testElementProcesserImpl{}
	p := Person{
		ID:           "24901",
		Name:         "steve_brine",
		Constituency: "winchester",
	}
	err := scrape(&ep, p)

	require.NoError(t, err)

	for _, t := range sections {
		fmt.Printf(" sections[\"%v\"]=[]string{}\n", t)
	}

	fmt.Printf(" === TOTAL: %v ===\n", len(sections))

	require.Equal(t, 93, len(sections), "there should be 93 sections voted on")
	require.True(t, false, "killed")
}

type testElementProcesserImpl struct {
}

// processVotePrintout - just a test func to print out the results
func (tep *testElementProcesserImpl) ProcessElement(index int, element *goquery.Selection) {
	fmt.Printf("VoteDesc: %v, %v\n", index, element.Text())
	element.Children().Each(processVoteChildrenPrintout)
}

func processVoteChildrenPrintout(index int, element *goquery.Selection) {
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

// TestScraperReal is how we should use the scraper to pull out someones record
func TestScraperReal(t *testing.T) {
	ep := ElementProcesserImpl{results: make(map[string]string)}
	p := Person{
		ID:           "24901",
		Name:         "steve_brine",
		Constituency: "winchester",
	}
	err := scrape(&ep, p)

	require.NoError(t, err)

	fmt.Println("==== RESULTS ====")
	for k, v := range ep.results {
		fmt.Printf("  [%v]:[%v]\n", k, v)
	}
	require.True(t, len(ep.results) > 0, "we should have found some results!")

	require.True(t, false, "killed")
}
