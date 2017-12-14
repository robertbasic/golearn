package lessons

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

		fmt.Println(fmt.Sprintf("%d\t%#U\t%t", b, b, b == '\n'))
	}
}

// ReadLineByLine reads a multiline string
// line by line
func ReadLineByLine() {
	br := bufio.NewReader(strings.NewReader(lines))

	for {
		l, err := br.ReadBytes('\n')

		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}

		fmt.Println(l)

		if err == io.EOF {
			break
		}
	}
}

// ReadLineByLineAsString reads a multiline string
// line by line as a string
func ReadLineByLineAsString() {
	br := bufio.NewReader(strings.NewReader(lines))

	for {
		// Includes the delimiter
		l, err := br.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}

		// Trimming space to remove the delimiter at the end
		fmt.Println(strings.TrimSpace(l))

		if err == io.EOF {
			break
		}
	}
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

	fmt.Println(lines)
}
