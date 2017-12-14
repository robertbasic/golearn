package lessons

import "testing"
import "log"
import "io/ioutil"

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
