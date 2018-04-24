package lib

type getCostInput struct {
	columnIdx   int
	previousRow []int
	currentRow  []int
	searchTerm  string
	letter      rune
}

// Get the minimum value of a slice of integers
func minInt(arr []int) int {
	min := arr[0]

	for _, v := range arr {
		if min > v {
			min = v
		}
	}

	return min
}

// Determine the distance between two points using the levenshtein algorithm
func getCost(input getCostInput) int {
	insertCost := input.currentRow[input.columnIdx-1] + 1
	deleteCost := input.previousRow[input.columnIdx] + 1
	var replaceCost int

	if []rune(input.searchTerm)[input.columnIdx-1] == input.letter {
		replaceCost = input.previousRow[input.columnIdx-1]
	} else {
		replaceCost = input.previousRow[input.columnIdx-1] + 1
	}

	return minInt([]int{insertCost, deleteCost, replaceCost})
}

type getResultsInput struct {
	node        *WordTree
	letter      rune
	previousRow []int
	searchTerm  string
	distance    int
}

// Get search results in the given word tree for the given search term
func getResults(input getResultsInput) []SearchResult {
	currentRow := []int{input.previousRow[0] + 1}
	var results []SearchResult

	// build out the row for the given letter
	columns := len(input.searchTerm) + 1
	for idx := 1; idx < columns; idx++ {
		currentRow = append(currentRow, getCost(getCostInput{
			columnIdx:   idx,
			previousRow: input.previousRow,
			currentRow:  currentRow,
			searchTerm:  input.searchTerm,
			letter:      input.letter,
		}))
	}

	// this word is the correct distance away, add it to the results slice
	if currentRow[len(currentRow)-1] <= input.distance && len(input.node.Word) > 0 {
		results = append(results, SearchResult{
			Word:     input.node.Word,
			Distance: currentRow[len(currentRow)-1],
		})
	}

	// if any items in the row are lower than the max distance, run the code again
	if minInt(currentRow) <= input.distance {
		for k, v := range input.node.Children {
			newResults := getResults(getResultsInput{
				node:        v,
				letter:      k,
				previousRow: currentRow,
				searchTerm:  input.searchTerm,
				distance:    input.distance,
			})

			results = append(results, newResults...)
		}
	}

	return results
}

// SearchResult is a struct which will be returned in a slice after
// a search
type SearchResult struct {
	Word     string
	Distance int
}

// WordSearch is a struct which will take a wordtree and provide
// a method to search it
type WordSearch struct {
	tree *WordTree
}

// NewWordSearch allocates and returns a new *WordTrie.
func NewWordSearch(tree *WordTree) *WordSearch {
	return &WordSearch{
		tree: tree,
	}
}

// Search using a search term within the given tree and return each
// item within a given distance of the search term
func (search *WordSearch) Search(searchTerm string, distance int) []SearchResult {
	currentRow := make([]int, len(searchTerm)+1)
	var results []SearchResult

	// Initialize the slice with the values equal to the keys
	for k := range currentRow {
		currentRow[k] = k
	}

	// Loop through each child of the word tree and concat search results
	for k, v := range search.tree.Children {
		newResults := getResults(getResultsInput{
			node:        v,
			letter:      k,
			previousRow: currentRow,
			searchTerm:  searchTerm,
			distance:    distance,
		})

		results = append(results, newResults...)
	}

	return results
}
