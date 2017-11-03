package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestInputOutput(t *testing.T) {
	sample := `First todo
Second todo
Third todo
`
	todo, err := readTodo(strings.NewReader(sample))
	if err != nil {
		t.Errorf("unexpected error %q", err)
	}

	var buf bytes.Buffer
	err = printTodo(todo, &buf)
	if err != nil {
		t.Errorf("unexpected error %q", err)
	}

	if buf.String() != sample {
		t.Errorf("expected %q got %q", sample, buf.String())
	}
}
