package eventbus

type Subscriber interface {
	OnEvent(evt string, data interface{})
}
