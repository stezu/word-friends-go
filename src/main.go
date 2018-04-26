package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/stezu/word-friends-go/src/lib"
)

const distance = 1

type networkResult struct {
	word    string
	results []lib.SearchResult
}

// write the number of friends since that's all the output we need
func writeResults(network []networkResult) {
	for _, v := range network {
		fmt.Println(len(v.results))
	}
}

func main() {
	networkUndefined := true
	tree := lib.NewWordTree()
	var network []networkResult

	scanner := bufio.NewScanner(os.Stdin)

	// Loop through each line of stdin and react to the three types of inputs
	// we can have in a file.
	for scanner.Scan() {
		text := scanner.Text()

		if text == "END OF INPUT" {
			networkUndefined = false
		} else if networkUndefined {
			network = append(network, networkResult{
				word:    text,
				results: nil,
			})
		} else {
			tree.Insert(text)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	wordSearch := lib.NewWordSearch(tree)

	// Search for friends of each word in the input
	for k, v := range network {
		searchResults := wordSearch.Search(v.word, distance)
		var filteredResults []lib.SearchResult

		// Remove all search results which are not exactly the right distance away
		for _, result := range searchResults {
			if result.Distance == distance {
				filteredResults = append(filteredResults, result)
			}
		}

		network[k].results = filteredResults
	}

	writeResults(network)
}
