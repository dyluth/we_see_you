package twitterscan

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanner(t *testing.T) {
	scanner, err := Init("../creds.json")
	require.NoError(t, err)
	require.NotZero(t, scanner)

	searchResult, _ := scanner.api.GetSearch("golang", nil)
	for _, tweet := range searchResult.Statuses {
		fmt.Printf("user: %v, says: %v\n", tweet.User.Name, tweet.Text)
	}

	require.True(t, false, "killed")
}
