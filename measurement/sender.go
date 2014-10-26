package measurement

const (
	ENDPOINT string = "www.google-analytics.com/collect"
)

// job queue
type stack struct {
	messages []string
}

// worker
type Worker struct {
}

func (w *Worker) Send(message string) {

}

// worker manager
type dispatcher struct {
	numWorkers int
	workers    []Worker
	queue      *stack
	comm       *chan string
}

func Dispatcher(workers int) *dispatcher {
	service := new(dispatcher)
	service.numWorkers = workers
	service.workers = make([]Worker, workers)
	service.queue = new(stack)
	service.comm = new(chan string)

	return service
}
