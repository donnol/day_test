package main

import "fmt"

var (
	MaxWorker = 100
	MaxQueue  = 100
)

// Job represents the job to be run
// 任务
type Job struct {
	Payload Payload
}

// A buffered channel that we can send work requests on.
// 任务管道
var JobQueue chan Job

// Worker represents the worker that executes the job
// 工作者
type Worker struct {
	WorkerPool chan chan Job // 管道池
	JobChannel chan Job      // 工作管道
	quit       chan bool     // 停止管道
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel // 缓存任务，将chan保存到池里面去

			select {
			case job := <-w.JobChannel: // 读取chan里面的值
				// we have received a work request.
				if err := job.Payload.UploadToS3(); err != nil {
					fmt.Printf("Error uploading to S3: %s", err.Error())
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

// 将任务发送到管道中
func payloadHandler3(data PayloadCollection) {

	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range data.Payloads {

		// let's create a job with the payload
		work := Job{Payload: payload}

		// Push the work onto the queue.
		JobQueue <- work
	}
}

// 分配／调度者
type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < MaxWorker; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

// 处理管道中的任务
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			// 一个任务开启一个goroutine处理
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool // 选择一个工作者

				// dispatch the job to the worker job channel
				jobChannel <- job // 将任务发送给工作者
			}(job)
		}
	}
}
