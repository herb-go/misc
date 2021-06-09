package busevent

type Event []*Handler

func (e *Event) Raise(data interface{}) {
	for _, v := range *e {
		if v.Fn != nil {
			v.Fn(data)
		}
	}
}

func (e *Event) Bind(l *Handler) {
	for _, v := range *e {
		if v.Key == l.Key {
			v.Fn = l.Fn
			return
		}
	}
	*e = append(*e, l)
}
func (l *Event) Unbind(key interface{}) {
	for _, v := range *l {
		if v.Key == key {
			v.Fn = nil
			return
		}
	}
}
func New() *Event {
	return &Event{}
}
