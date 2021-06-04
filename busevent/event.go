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
		if v.ID == l.ID {
			v.Fn = l.Fn
			return
		}
	}
	*e = append(*e, l)
}
func (l *Event) Unbind(id interface{}) {
	for _, v := range *l {
		if v.ID == id {
			v.Fn = nil
			return
		}
	}
}
func New() *Event {
	return &Event{}
}
