package lessons

import (
	"bufio"
	"fmt"
	"io"
)

// ReadTodo reads in Inputs from the given reader,
// until it gets an empty line
// and puts them into a slice of strings
func ReadTodo(input io.Reader) ([]string, error) {
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

// PrintTodo prints the slice of todos
func PrintTodo(todo []string, output io.Writer) error {
	for _, v := range todo {
		_, err := fmt.Fprintln(output, v)
		if err != nil {
			return fmt.Errorf("failed at printing to output: %s", err)
		}
	}
	return nil
}
