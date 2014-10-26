package measurement

import (
	"testing"
)

func TestInit(t *testing.T) {
	numWorkers := 5
	dispatcher := Dispatcher(numWorkers)
	if dispatcher.numWorkers != numWorkers || len(dispatcher.workers) != numWorkers {
		t.Error("Failed initiating dispatcher")
	}
}
