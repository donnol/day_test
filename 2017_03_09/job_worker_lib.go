// 模拟工作／工作者模式
// 通过环境变量指定工作者数量、队列数量
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// 工作
type Job struct {
	Data    interface{}                  // 工作要处理的内容
	Handler func(data interface{}) error // 工作处理方法
}

// 工作管道
var jobQueue chan Job

// 工作者
type Worker struct {
	workerPool chan chan Job // 保存工作管道的管道
	jobChannel chan Job      // 工作管道
	quit       chan bool     // 停止管道
}

// 创建工作者
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		workerPool: workerPool,
		jobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// 启动工作者
func (w Worker) Start() {
	go func() {
		for {
			// 保存工作管道到工作池管道中
			w.workerPool <- w.jobChannel

			select {
			// 读取JobChannel的值
			case job := <-w.jobChannel:
				// 打印工作名
				err := job.Handler(job.Data)
				if err != nil {
					log.Println(err)
				}
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
	WorkerPool chan chan Job // 工作管道池
}

// 创建调度器
func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Job, maxWorkers),
	}
}

// 运行调度器
func (d *Dispatcher) Run() {
	// 启动指定数量的工作者
	for i := 0; i < MaxWorker; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

// 调度处理
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-jobQueue:
			// 接收到工作
			go func(job Job) {
				jobChannel := <-d.WorkerPool

				jobChannel <- job
			}(job)
		}
	}
}

func AddJob(names []string) {
	for _, single := range names {
		// 创建工作
		work := Job{
			Data: single,
			Handler: func(data interface{}) error {
				if url, ok := data.(string); ok {
					resp, err := http.Get(url)
					if err != nil {
						return err
					}
					defer resp.Body.Close()

					content, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						return err
					}
					err = ioutil.WriteFile("model.html", content, 0666)
					if err != nil {
						return err
					}
				}
				return nil
			},
		}

		// 将工作推入工作管道
		JobQueue <- work
	}
}
