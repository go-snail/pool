package pool

import (
	"sync"
)


/* 有关协程池的定义及操作 */
//定义池类型
type Pool struct {

	//对外接收Task的入口
	EntryChannel chan *Message

	//分发器
	 IDispatcher

	//协程池最大worker数量,限定Goroutine的个数
	taskQueueNum int

	//TaskQueue队列
	TQ[] *TaskQueue

	//Worker处理器
	Worker *Worker
	// 等待任务池中的协程启动完成
	sync.WaitGroup
}

func NewPool(num int) *Pool {
	entryChannel := make(chan *Message)
	dp := NewDispatcher()
	taskQueue := make([]*TaskQueue, num)
	for i := range taskQueue {
		taskQueue[i] = NewTaskQueue(i)
	}
	worker := NewWorker()

	return &Pool{EntryChannel: entryChannel,IDispatcher:dp, taskQueueNum: num, TQ: taskQueue, Worker: worker}
}

// 用户可自定义dispatcher调度器
// 只需要实现IDispatcher接口即可
func (p *Pool)SetDispatcher(dp IDispatcher)  {
	p.IDispatcher = dp
}


func (p *Pool)Run()  {
	p.Add(p.taskQueueNum)
	for i := 0; i < p.taskQueueNum; i++ {
		p.Done()
		go p.Worker.do(p.TQ[i])
	}
	p.Wait()
	//todo 读取p.EntryChannel
	go func() {
		for {
			select {
			case m := <-p.EntryChannel:
				p.dispatcher(m,p)
			}
		}
	}()
}
