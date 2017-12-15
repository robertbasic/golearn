package lessons

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var lines = `first
second
third`

// ReadLinesByteByByte reads a multiline string
// byte by byte
func ReadLinesByteByByte() {
	r := strings.NewReader(lines)

	for {
		b, err := r.ReadByte()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		log.Println(fmt.Sprintf("%d\t%#U\t%t", b, b, b == '\n'))
	}
}

// ReadLineByLine reads a multiline string
// line by line
func ReadLineByLine() {
	br := bufio.NewReader(strings.NewReader(lines))

	var lines []string

	for {
		l, err := br.ReadBytes('\n')

		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}

		lines = append(lines, strings.TrimSpace(string(l)))

		if err == io.EOF {
			break
		}
	}

	log.Println(lines)
}

// ReadLineByLineAsString reads a multiline string
// line by line as a string
func ReadLineByLineAsString() {
	br := bufio.NewReader(strings.NewReader(lines))

	var lines []string

	for {
		// Includes the delimiter
		l, err := br.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}

		// Trimming space to remove the delimiter at the end
		lines = append(lines, strings.TrimSpace(l))

		if err == io.EOF {
			break
		}
	}

	log.Println(lines)
}

// ReadLineByLineWithAnonFunc reads a multiline string
// line by line as a string, but uses an anon func to
// do the actual reading.
func ReadLineByLineWithAnonFunc() {
	br := bufio.NewReader(strings.NewReader(lines))

	var lines []string

	readLine := func() (string, bool) {
		l, err := br.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}

		return strings.TrimSpace(l), err == io.EOF
	}
	for {
		l, e := readLine()
		lines = append(lines, l)

		if e {
			break
		}
	}

	log.Println(lines)
}

// ReadFile reads a file
func ReadFile() {
	file, err := getFile()
	if err != nil {
		log.Fatalln(err)
	}

	br := bufio.NewReader(file)

	var lines []string

	for {
		// Includes the delimiter
		l, err := br.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}

		// Trimming space to remove the delimiter at the end
		lines = append(lines, strings.TrimSpace(l))

		if err == io.EOF {
			break
		}
	}

	log.Println(lines)
}

func getFile() (io.Reader, error) {
	wd, _ := os.Getwd()

	if !strings.HasSuffix(wd, "lessons") {
		wd += "/lessons"
	}

	file, err := os.Open(wd + "/biteme.txt")

	if err != nil {
		return nil, err
	}

	return file, nil
}
