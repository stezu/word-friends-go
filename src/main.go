package main

import (
	"bufio"
	"os"

	"github.com/stezu/word-friends-go/src/lib"
)

const distance = 1

func getFileFromStdin() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

func getFileFromFilename(fileName string) (*bufio.Scanner, error) {
	file, err := os.Open(os.Args[1])

	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(file), nil
}

func parseArgs() (*bufio.Scanner, error) {
	// If the first argument exists, that means we are reading from a file
	if len(os.Args) > 1 {
		fileScanner, err := getFileFromFilename(os.Args[1])

		if err != nil {
			return nil, err
		}

		return fileScanner, nil
	}

	// Otherwise, we are reading from stdin
	return getFileFromStdin(), nil
}

func buildNetwork(scanner *bufio.Scanner) (*lib.Network, error) {
	network := lib.NewNetwork()
	networkDefined := false

	// Loop through each line of the stream and react to the three types of inputs
	// we can have in a file.
	for scanner.Scan() {
		text := scanner.Text()

		if text == "END OF INPUT" {
			networkDefined = true
		} else if !networkDefined {
			network.AddWord(text)
		} else {
			network.Search(text, distance)
		}
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return network, nil
}

func main() {
	scanner, err := parseArgs()

	if err != nil {
		panic(err)
	}

	network, err := buildNetwork(scanner)

	if err != nil {
		panic(err)
	}

	network.PrintResults(distance)
}
