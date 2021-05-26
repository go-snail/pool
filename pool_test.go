package pool

import (
	"fmt"
	"runtime"
	"testing"
)

const poolNum = 3

// 自定义message处理method
func doMessage(message *Message)  {
	fmt.Println(",doMessage:msgId:",message.MsgId,"msgType:",string(message.Data))
}



func TestPool(t *testing.T) {
	// 自定义worker，然后注册到pool池中，pool池回掉worker处理消息



}

func BenchmarkPool(b *testing.B)  {
	cpuNum:=runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
	p := NewPool(poolNum)
	//p.Worker.SetFunc(doMessage)
	p.Run()
	for i := 0; i < b.N; i++ {
		//构造task,向EntryChannel发送
		msg := &Message{i,1,[]byte("hello world!")}
		p.EntryChannel <- msg
	}
}