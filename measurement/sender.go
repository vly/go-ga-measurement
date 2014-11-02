package measurement

const (
	ENDPOINT string = "www.google-analytics.com/collect"
)

// job queue
type stack struct {
	messages []string
}

func (s stack) Empty() bool {
	return len(s.messages) == 0
}

func (s stack) Peek() string {
	return s.messages[len(s.messages)-1]
}

func (s stack) Len() string {
	return len(s.messages)
}

func (s *stack) Put(i string) {
	s.messages = append(s.messages, i)
}

func (s *stack) Pop() string {
	mesg := s.messages[len(s.messages)-1]
	s.messages = s.messages[:len(s.messages)-1]
	return mesg
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
