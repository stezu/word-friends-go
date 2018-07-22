package lib

import (
	"fmt"

	"github.com/arbovm/levenshtein"
)

type searchResult struct {
	distance int
	word     string
}

func newSearchResult(word string, distance int) *searchResult {
	return &searchResult{
		distance: distance,
		word:     word,
	}
}

type networkEntry struct {
	matches []searchResult
	word    string
}

func newNetworkEntry(word string) *networkEntry {
	return &networkEntry{
		matches: make([]searchResult, 0),
		word:    word,
	}
}

func (entry *networkEntry) addMatch(result searchResult) *networkEntry {
	entry.matches = append(entry.matches, result)

	return entry
}

func (entry *networkEntry) search(word string, maxDistance int) *networkEntry {
	actualDistance := levenshtein.Distance(entry.word, word)

	if actualDistance <= maxDistance {
		return entry.addMatch(*newSearchResult(word, actualDistance))
	}

	return entry
}

// Network a network of words
type Network struct {
	entries []networkEntry
}

// NewNetwork allocates and returns a new *Network.
func NewNetwork() *Network {
	return &Network{
		entries: make([]networkEntry, 0),
	}
}

// AddWord adds a word to the network
func (network *Network) AddWord(word string) *Network {
	network.entries = append(network.entries, *newNetworkEntry(word))

	return network
}

// Search finds the matching relatives in the entire network
func (network *Network) Search(word string, distance int) *Network {
	for idx, entry := range network.entries {
		network.entries[idx] = *entry.search(word, distance)
	}

	return network
}

// PrintResults will print the number of matches which have the given distance from each word in the network
func (network *Network) PrintResults(distance int) *Network {
	for _, entry := range network.entries {
		count := 0

		for _, match := range entry.matches {
			if match.distance == distance {
				count++
			}
		}

		fmt.Println(count)
	}

	return network
}
