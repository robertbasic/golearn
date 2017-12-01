package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	todo, err := readTodo(os.Stdin)
	if err != nil {
		log.Fatalf("failed at reading todo: %s", err)
	}

	fmt.Printf("Got %d inputs from stdin:\n", len(todo))

	err = printTodo(todo, os.Stdout)
	if err != nil {
		log.Fatalf("failed at printing todo: %s", err)
	}
}

// readTodo reads in Inputs from the given reader,
// until it gets an empty line
// and puts them into a slice of strings
func readTodo(input io.Reader) ([]string, error) {
	var todo []string

	scanner := bufio.NewScanner(input)

	askAndScan := func() bool {
		fmt.Print("Add a new todo [empty to stop reading]: ")
		return scanner.Scan()
	}
	for askAndScan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		todo = append(todo, line)
	}
	if err := scanner.Err(); err != nil {
		return todo, fmt.Errorf("failed at scanning: %s", err)
	}

	return todo, nil
}

func printTodo(todo []string, output io.Writer) error {
	for _, v := range todo {
		_, err := fmt.Fprintln(output, v)
		if err != nil {
			return fmt.Errorf("failed at printing to output: %s", err)
		}
	}
	return nil
}
