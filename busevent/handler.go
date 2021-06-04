package busevent

type Handler struct {
	ID interface{}
	Fn func(data interface{})
}

func NewHandler() *Handler {
	return &Handler{}
}

func CreateHandler(id interface{}, fn func(data interface{})) *Handler {
	return &Handler{
		ID: id,
		Fn: fn,
	}
}
