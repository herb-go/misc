package busevent

type Listener struct {
	Key     interface{}
	Handler func(data interface{})
}

func (l *Listener) WithKey(key interface{}) *Listener {
	l.Key = key
	return l
}
func (l *Listener) WithHandler(handler func(data interface{})) *Listener {
	l.Handler = handler
	return l
}
func NewListener() *Listener {
	return &Listener{}
}
