package lib

import "fmt"

// WordTree a tree of words
type WordTree struct {
	Children map[rune]*WordTree
	Word     string
}

// NewWordTree allocates and returns a new *WordTree.
func NewWordTree() *WordTree {
	return &WordTree{
		Children: make(map[rune]*WordTree),
	}
}

// Insert a word into the tree, optionally branch off for every new letter
func (tree *WordTree) Insert(word string) {
	node := tree

	for _, r := range word {
		child, _ := node.Children[r]

		if child == nil {
			child = NewWordTree()
			node.Children[r] = child
		}

		node = child
	}

	node.Word = word
}

// WordCount get the number of words in this tree
func (tree *WordTree) WordCount() int {
	var count int

	if len(tree.Word) > 0 {
		count++
	}

	fmt.Println(count)

	for _, subTree := range tree.Children {
		count += subTree.WordCount()
	}

	return count
}

// NodeCount get the number of nodes in this tree
func (tree *WordTree) NodeCount() int {
	count := 1

	for _, subTree := range tree.Children {
		count += subTree.NodeCount()
	}

	return count
}

// WalkFunc defines some action to take on the given key and value during
// a Tree Walk. Returning a non-nil error will terminate the Walk.
// type WalkFunc func(key string, value string) error

// Walk iterates over each key/value stored in the tree and calls the given
// walker function with the key and value. If the walker function returns
// an error, the walk is aborted.
// func (tree *WordTree) Walk(walker WalkFunc) error {
// 	return tree.walk("", walker)
// }

// func (tree *WordTree) walk(key string, walker WalkFunc) error {
// 	if len(tree.Word) > 0 {
// 		walker(key, tree.Word)
// 	}

// 	for r, child := range tree.Children {
// 		err := child.walk(key+string(r), walker)

// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
