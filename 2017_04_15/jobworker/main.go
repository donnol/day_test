// 模拟工作／工作者模式
// 通过环境变量指定工作者数量、队列数量
package jobworker

// 工作--外部添加的工作
type job struct {
	data    interface{}            // 工作要处理的内容
	handler func(data interface{}) // 工作处理方法
}

// 工作管道
var jobQueue chan job

// 工作者
type Worker struct {
	workerPool chan chan job // 保存工作管道的管道
	jobChannel chan job      // 工作管道
	quit       chan bool     // 停止管道
}

// 创建工作者
func NewWorker(workerPool chan chan job) Worker {
	return Worker{
		workerPool: workerPool,
		jobChannel: make(chan job),
		quit:       make(chan bool),
	}
}

// 启动工作者
func (w Worker) Start() {
	go func() {
		for {
			// 保存工作管道到工作池管道中
			w.workerPool <- w.jobChannel // 在装满chan时，将阻塞，直到有位置空出

			select {
			// 读取JobChannel的值
			case job := <-w.jobChannel:
				// 处理
				job.handler(job.data)
			// 停止
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

// 调度器
type Dispatcher struct {
	WorkerPool chan chan job // 工作管道池
}

// 创建调度器
func NewDispatcher(maxWorker int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan job, maxWorker), // 初始化worker chan池
	}
}

// 运行调度器
func (d *Dispatcher) Run(maxWorker, maxQueue int) {
	jobQueue = make(chan job, maxQueue) // 初始化工作队列

	// 启动指定数量的worker
	for i := 0; i < maxWorker; i++ {
		worker := NewWorker(d.WorkerPool) // 初始化worker，所有worker公用一个worker chan池
		worker.Start()                    // 启动worker
	}

	go d.dispatch()
}

// 调度处理
// 调度器读取jobQueue的工作，并从worker chan池中选择一个worker chan去处理job
func (d *Dispatcher) dispatch() {
	for {
		select {
		case singleJob := <-jobQueue:
			// 接收到工作
			go func(singleJob job) {
				jobChannel := <-d.WorkerPool // 从worker chan池中取出一个worker chan去处理job

				jobChannel <- singleJob // 将job发到job chan
			}(singleJob)
		}
	}
}

// 添加工作
func CommonAddJob(data interface{}, handler func(data interface{})) {
	work := job{
		data:    data,
		handler: handler,
	}

	// 将工作推入工作管道
	jobQueue <- work
}
