package lessons

import (
	"log"
	"sync"
)

// RoutinesWaitGroup teaches about goroutines
// using sync.WaitGroup to wait for goroutines to finish
func RoutinesWaitGroup() {
	log.Println("Started.")

	var wg sync.WaitGroup
	wg.Add(2)
	go rwg("a", &wg)
	go rwg("b", &wg)
	wg.Wait()

	log.Println("Done.")
}

// RoutinesWaitChannel teaches about goroutines
// using channels to communicate when all is said and done
func RoutinesWaitChannel() {
	log.Println("Started.")

	// var donea chan bool doesn't work!
	donea := make(chan bool)
	doneb := make(chan bool)

	go rwc("a", donea)
	go rwc("b", doneb)

	<- donea
	<- doneb

	close(donea)
	close(doneb)

	log.Println("Done.")
}

// RoutineTalk teaches about communication between goroutines
// using channels
func RoutineTalk() {
	log.Println("Started.")

	msg := make(chan string)
	done := make(chan bool)

	go senderRoutine(msg)
	go receiverRoutine(msg, done)

	<- done

	close(msg)
	close(done)

	log.Println("Done.")
}

func senderRoutine(msg chan string) {
	msg <- "Hello from sender."
}

func receiverRoutine(msg chan string, done chan bool) {
	log.Println(<- msg)
	done <- true
}

func rwc(s string, done chan bool) {
	for i := 0; i < 10; i++ {
		log.Println(s, i)
	}
	done <- true
}

func rwg(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		log.Println(s, i)
	}
}