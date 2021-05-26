package pool

type IDispatcher interface {
	dispatcher(message *Message,p *Pool)
}
