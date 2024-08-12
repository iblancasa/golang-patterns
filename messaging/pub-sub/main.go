package main

import (
	"context"
	"log"
	"sync"
	"time"
)

type Message struct {
	m      string
	sender int
}

func (m Message) String() string {
	return m.m
}

type Subscriber struct {
	id int
	Ch chan Message
}

func newSubscriber(id int) Subscriber {
	return Subscriber{
		id: id,
		Ch: make(chan Message),
	}
}

func (s *Subscriber) listen(ctx context.Context) {
	channelOpen := true
	var m Message
	for channelOpen {
		select {
		case m, channelOpen = <-s.Ch:
			log.Printf("Message received in subscriber '%d' from '%d': %s\n", s.id, m.sender, m)
		case <-ctx.Done():
			log.Println("Stopping subscriber", s.id)
			return
		}
	}
}

type Publisher struct {
	id          int
	subscribers []*Subscriber
	mu          sync.Mutex
}

func (p *Publisher) send(m string) {
	p.mu.Lock()
	for _, s := range p.subscribers {
		s.Ch <- Message{m: m, sender: p.id}
	}
	p.mu.Unlock()
}

func (p *Publisher) Subscribe(s *Subscriber) {
	p.mu.Lock()
	p.subscribers = append(p.subscribers, s)
	p.mu.Unlock()
}

func main() {
	ctx := context.Background()
	wg := sync.WaitGroup{}
	p1 := Publisher{id: 1}

	for i := 0; i < 10; i++ {
		s := newSubscriber(i)
		p1.Subscribe(&s)
		go s.listen(ctx)
	}

	wg.Add(1)
	defer wg.Wait()
	go func() {
		for i := 0; i < 10000; i++ {
			p1.send("hello")
			time.Sleep(time.Millisecond)
		}
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		s := newSubscriber(i)
		p1.Subscribe(&s)
		go s.listen(ctx)
	}

}
