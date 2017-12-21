package lessons

import "log"

func RunGenerator() {
	c := make(chan int)

	go gen(c)

	log.Println(<-c)
}

func gen(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
}