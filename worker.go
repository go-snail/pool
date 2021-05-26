package pool

import (
	"fmt"
)

type Worker struct {
	f func(message *Message)
}

func defaultf(msg *Message)  {
	fmt.Println(",messageID:",msg.MsgId,",msg:",string(msg.Data))
}

func NewWorker() *Worker {
	worker := new(Worker)
	worker.f = defaultf
	return worker
}

func (w *Worker)SetFunc(f func(message *Message)) {
	w.f = f
}


func (w *Worker)do(taskQueue *TaskQueue)  {
	for {
		select {
		case msg := <-taskQueue.mchan:
			fmt.Printf("tid:%d",taskQueue.tid)
			w.f(msg)
		}
	}

}

