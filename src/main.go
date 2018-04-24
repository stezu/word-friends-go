package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/stezu/word-friends-go/src/lib"
)

const distance = 1

// write the number of friends since that's all the output we need
func writeResults(network map[string][]lib.SearchResult) {
	for _, v := range network {
		fmt.Println(len(v))
	}
}

func main() {
	networkUndefined := true
	network := make(map[string][]lib.SearchResult)
	tree := lib.NewWordTree()

	scanner := bufio.NewScanner(os.Stdin)

	// Loop through each line of stdin and react to the three types of inputs
	// we can have in a file.
	for scanner.Scan() {
		text := scanner.Text()

		if text == "END OF INPUT" {
			networkUndefined = false
		} else if networkUndefined {
			network[text] = []lib.SearchResult{} // TODO: think about this differently?
		} else {
			tree.Insert(text)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	wordSearch := lib.NewWordSearch(tree)

	// Search for friends of each word in the input
	for k := range network {
		searchResults := wordSearch.Search(k, distance)
		var filteredResults []lib.SearchResult

		// Remove all search results which are not exactly the right distance away
		for _, result := range searchResults {
			if result.Distance == distance {
				filteredResults = append(filteredResults, result)
			}
		}

		network[k] = filteredResults
	}

	writeResults(network)
}
