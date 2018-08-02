package analyser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDumbAnalyser(t *testing.T) {

	da := NewDumbAnalyser()
	require.True(t, MessageTest(&da, "this is my string about the LGBT community and how thats a good thing", "gay rights"))

	require.True(t, MessageTest(&da, "we have a smoking ban in enclosed public places for people's health", "smoking bans"))

	require.True(t, MessageTest(&da, "couples of the same sex should have the same rights and opportunities as everyone else", "marriage"))

	require.True(t, MessageTest(&da, "human rights violations should be condemned and action should be taken against those that do them", "equality and human rights"))

}

func MessageTest(a Analyser, message string, expectedCategory string) bool {

	cat, conf := a.GetCategory(message)
	if conf < 70 {
		fmt.Printf("confidence too low (%v)\n", conf)
		return false
	}
	if cat == expectedCategory {
		fmt.Printf("category wrong, wanted: [%v], have [%v]\n", expectedCategory, cat)
		return true
	}
	return false
}
