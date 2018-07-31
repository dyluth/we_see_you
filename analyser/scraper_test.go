package analyser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScraper(t *testing.T) {
	scrape()
	require.True(t, false, "killed")
}
