package main

import (
	"bufio"
	"fmt"
	"os"
)

// Inputs holds the inputs read in from stdin
type Inputs []string

// ReadInput reads in Inputs from the standard
// input, until it gets an empty line
// and puts them into a slice of strings, Inputs
func ReadInput(scanner bufio.Scanner) Inputs {
	var t string
	var i Inputs

	for true {
		fmt.Print("Add a new todo [empty to stop reading]: ")
		scanner.Scan()
		t = scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from input: ", err)
		}

		// scanner.Text() strips new lines
		// so in case of just a new line
		// it's actually an empty string
		if t == "" {
			break
		}

		i = append(i, t)
		fmt.Println(t)
	}

	return i
}

func main() {
	fmt.Println("Hello, world!")

	scanner := bufio.NewScanner(os.Stdin)
	is := ReadInput(*scanner)

	fmt.Printf("Got %d inputs from stdin:\n", len(is))

	for _, v := range is {
		fmt.Println(v)
	}
}
