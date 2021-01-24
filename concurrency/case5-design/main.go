// https://www.mofish.work/thread/8527
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Payload struct {
	name string
}

//任務
type Job struct {
	Payload Payload
}

//任務佇列
var JobQueue chan Job

//  執行者 消費者 工人
type Worker struct {
	WorkerPool chan chan Job //物件池
	JobChannel chan Job      //通道里面拿
	quit       chan bool     //
	name       string        //工人的名字
}

// 排程器
type Dispatcher struct {
	name       string        //排程的名字
	maxWorkers int           //獲取 除錯的大小
	WorkerPool chan chan Job //註冊和工人一樣的通道
}

//打遊戲
func (p *Payload) Play() {
	fmt.Printf("%s 打 LOL 遊戲...當前任務完成\n", p.name)
}

// 新建一個工人
func NewWorker(workerPool chan chan Job, name string) Worker {
	fmt.Printf("建立了一個工人,它的名字是:%s \n", name)
	//workerPool 確定 woker 的容量
	return Worker{
		name:       name,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

// 工人開始工作

func (w *Worker) Start() {
	//開一個新的協程
	go func() {
		for {
			//註冊到物件池中,
			// jobChannel := <-d.WorkerPool
			// jobChannel <- job //其实是 w.JobChannel <- job

			// 所以
			// select {
			// //会接受到新流入的 channel
			// case job := <-w.JobChannel:
			w.WorkerPool <- w.JobChannel
			fmt.Printf("[%s]把自己註冊到 物件池中 \n", w.name)
			select {
			//接收到了新的任務
			case job := <-w.JobChannel:
				fmt.Printf("[%s] 工人接收到了任務 當前任務的長度是[%d]\n", w.name, len(w.WorkerPool))
				job.Payload.Play()
				time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			//接收到了任務
			case <-w.quit:
				fmt.Println("結束任務", w.name)
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

func NewDispatcher(maxWorkers int) *Dispatcher {
	//容量為{maxWorkers}的 channel
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,       // 將工人放到一個池中,可以理解成一個部門中
		name:       "排程者",      //排程者的名字
		maxWorkers: maxWorkers, //這個排程者有好多個工人
	}
}

func (d *Dispatcher) Run() {
	// 開始執行
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, fmt.Sprintf("work-%s", strconv.Itoa(i)))
		//開始工作
		worker.Start()
	}
	//監控
	go d.dispatch()

}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("排程者,接收到一個工作任務")
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			// 排程者接收到一個工作任務
			go func(job Job) {
				//從現有的物件池中拿出一個
				jobChannel := <-d.WorkerPool
				fmt.Println(jobChannel)
				fmt.Println(job)
				jobChannel <- job
			}(job)
		default:

			//fmt.Println("ok!!")
		}

	}
}

func initialize() {
	maxWorkers := 1
	maxQueue := 20
	//初始化一個除錯者,並指定它可以操作的 工人個數
	dispatch := NewDispatcher(maxWorkers)
	JobQueue = make(chan Job, maxQueue) //指定任務的佇列長度
	//並讓它一直接執行
	dispatch.Run()
}

func main() {
	//初始化物件池
	initialize()
	for i := 0; i < 2; i++ {
		p := Payload{
			fmt.Sprintf("玩家-[%s]", strconv.Itoa(i)),
		}
		JobQueue <- Job{
			Payload: p,
		}
		time.Sleep(time.Second)
	}
	close(JobQueue)
}
