package lessons

import (
	"log"
	"sync"
)

type message struct {
	sender int
	val int
}

type Program struct {
	id  int
	sch chan message
	rch chan message
	count int
}

func newProgram(id int) *Program {
	sch := make(chan message)
	p := Program{
		id:  id,
		sch: sch,
	}
	return &p
}

func TalkingGos() {
	var wg sync.WaitGroup
	wg.Add(2)

	p0 := newProgram(0)
	p1 := newProgram(1)

	p0.rch = p1.sch
	p1.rch = p0.sch

	go run(p0, &wg)
	go run(p1, &wg)

	wg.Wait()

	close(p0.sch)
	close(p1.sch)
}

func run(p *Program, wg *sync.WaitGroup) {
	var send = func(i int) {
		m := message{
			sender: p.id,
			val: i,
		}
		log.Println("Sending message from: ", p.id)
		p.sch <- m
	}
	var receive = func() {
		for {
			select {
			case m := <-p.rch:
				log.Printf("Got message, sender %d, val %d\n", m.sender, m.val)

				if m.val == 9 {
					wg.Done()
				}
			}
		}
	}
	for i := 0; i < 10; i++ {
		go send(i)
		go receive()
	}
}
