package main

import "log"
import "os"
import "os/signal"
import "syscall"
import "net"
import "time"
import "runtime"

var queue = &Queue{make(chan *Story, 1000)}

func main() {
	listener, err := net.Listen("unix", "/tmp/story_teller.sock")

	if err != nil {
		log.Fatal("Couldn't launch agent: ", err)
	}

	go theEnd(listener)

	go func() {
		clock := time.Tick(1 * time.Second)
		var m runtime.MemStats
		for range clock {
			stories := queue.Collect()
			runtime.ReadMemStats(&m)
			log.Print("Messages that would be processed: ", len(stories))
		}
	}()

	for {
		process(listener.Accept())
	}
}

func process(c net.Conn, err error) (*Story, error) {
	if err != nil {
		log.Print("Connection error: ", err)
	}

	defer c.Close()

	buf := make([]byte, 2048)
	nr, err := c.Read(buf)

	if err != nil {
		return nil, err
	}

	story, err := NewStoryFromJSON(buf[:nr])

	if err != nil {
		log.Print("Couldn't parse the following Story: ", string(buf), err)
		return nil, err
	}

	// log.Print("Message: ", story)

	queue.Add(story)

	return story, nil
}

func theEnd(l net.Listener) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)

	<-channel
	l.Close()
	os.Exit(0)
}
