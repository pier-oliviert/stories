package main

type Queue struct {
	channel chan *Story
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

func (q *Queue) Add(story *Story) {
	q.channel <- story
}
