package main

import (
	"fmt"
	"log"
	"os"

	"github.com/robertbasic/golearn/lessons"
)

func main() {
	lessons.ReadLinesByteByByte()
	lessons.ReadLineByLine()
	lessons.ReadLineByLineAsString()
	lessons.ReadLineByLineWithAnonFunc()
	lessons.ReadLinesWithScanner()
	lessons.ReadFile()
	lessons.ReadFileWithScanner()

	// readTodo()
}

func readTodo() {
	todo, err := lessons.ReadTodo(os.Stdin)
	if err != nil {
		log.Fatalf("failed at reading todo: %s", err)
	}

	fmt.Printf("Got %d inputs from stdin:\n", len(todo))

	err = lessons.PrintTodo(todo, os.Stdout)
	if err != nil {
		log.Fatalf("failed at printing todo: %s", err)
	}
}
