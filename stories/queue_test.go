package stories

import (
	"testing"
)

func TestQueueInitializedWithSize(t *testing.T) {
	size := 100
	queue := NewQueueOfSize(size)

	if queue.Size() != size {
		t.Fail()
	}
}

func TestCollectingWillClearTheQueue(t *testing.T) {
}

func TestQueueIsFullAtMaxSize(t *testing.T) {
}

func TestEmptyQueue(t *testing.T) {
}
