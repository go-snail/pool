package pool
const (
	TaskQueneLen = 1024
)
/* 有关Task任务相关定义及操作 */
//定义任务Task类型,每一个任务Task都可以抽象成一个函数
type TaskQueue struct {
	tid int
	mchan chan *Message
}

func NewTaskQueue(id int) *TaskQueue {
	mc := make(chan *Message,TaskQueneLen)
	return &TaskQueue{tid:id, mchan: mc}
}


