package lessons

import (
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkReadLineByLine(b *testing.B) {

	log.SetOutput(ioutil.Discard)

	for n := 0; n < b.N; n++ {
		ReadLineByLine()
	}
}
func BenchmarkReadLineByLineAsString(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	for n := 0; n < b.N; n++ {
		ReadLineByLineAsString()
	}
}

func BenchmarkReadLineByLineWithAnonFunc(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	for n := 0; n < b.N; n++ {
		ReadLineByLineWithAnonFunc()
	}
}

func BenchmarkReadFile(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	for n := 0; n < b.N; n++ {
		ReadFile()
	}
}

func BenchmarkReadFileWithScanner(b *testing.B) {
	log.SetOutput(ioutil.Discard)

	for n := 0; n < b.N; n++ {
		ReadFileWithScanner()
	}
}
