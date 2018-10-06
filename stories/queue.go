package stories

import (
	"log"
	"net"
)

type Analytics interface {
	PrintStats() string
}

type Queue struct {
	channel chan *Story
	*Stats
}

type Stats struct {
}

func NewQueueOfSize(size int) *Queue {
	log.Printf("Starting the queue with a buffer of %d stories", size)
	return &Queue{make(chan *Story, size), &Stats{}}
}

func (q *Queue) Add(c net.Conn) error {
	defer c.Close()

	buf, nr, err := read(c)

	if err != nil {
		return err
	}

	story, err := NewStory(buf[:nr])

	if err != nil {
		return err
	}

	q.channel <- story

	return nil
}

func (q *Queue) Collect() []*Story {
	var stories []*Story
	empty := false

	for empty != true {
		var story *Story
		select {
		case story = <-q.channel:
			stories = append(stories, story)
		default:
			empty = true
		}
	}

	return stories
}

func (q *Queue) Size() int {
	return cap(q.channel)
}

func (q *Queue) InQueue() int {
	return len(q.channel)
}

func (q *Queue) IsEmpty() bool {
	return len(q.channel) == 0
}

func (q *Queue) IsFull() bool {
	return q.InQueue() == q.Size()
}

func read(c net.Conn) ([]byte, int, error) {
	buf := make([]byte, 2048)
	nr, err := c.Read(buf)

	return buf, nr, err
}
