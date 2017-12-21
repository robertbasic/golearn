package lessons

import (
	"log"
	"fmt"
)

func RunGenerator() {
	for i := range numgenc(100) {
		log.Println(fmt.Sprintf("%T", i))
		log.Println(i)
	}
}

func numgenc(l int) chan int {
	c := make(chan int)

	go func() {
		defer close(c)

		for i := 0; i < l; i++ {
			c <- i*2
		}
	}()

	return c
}
