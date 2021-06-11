package busevent

type Event []*Listener

func (e *Event) Raise(data interface{}) {
	for _, v := range *e {
		v.Handler(data)
	}
}
func (e *Event) Bind(handler func(data interface{})) {
	*e = append(*e, NewListener().WithHandler(handler))
}
func (e *Event) BindAs(key interface{}, handler func(data interface{})) {
	*e = append(*e, NewListener().WithKey(key).WithHandler(handler))
}

func (e *Event) Remove(key interface{}) {
	if key == nil {
		return
	}
	var result = Event{}
	for _, v := range *e {
		if v.Key != key {
			result = append(result, v)
		}
	}
	*e = result
}
func (e *Event) Flush() {
	*e = nil
}

func New() *Event {
	return &Event{}
}
