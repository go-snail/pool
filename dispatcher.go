package pool

type DisPatcher struct {

}

func NewDispatcher()  *DisPatcher {
	return new(DisPatcher)
}

// 默认的任务调度器，基于messageID分发
func (d *DisPatcher)dispatcher(m *Message, p *Pool)  {
	tqID := m.MsgId % p.taskQueueNum
	p.TQ[tqID].mchan <- m
}
