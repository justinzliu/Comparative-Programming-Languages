package work_queue

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	q.Jobs = make(chan Worker, maxJobs)
	q.Results = make(chan interface{})
	for i := uint(0); i < nWorkers; i++ {
		go func(){
			q.worker()
		}()
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	for i := range queue.Jobs {
		task := i
		result := task.Run()
		queue.Results <- result
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	close(queue.Jobs)
}
