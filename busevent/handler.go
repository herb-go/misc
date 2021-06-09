package busevent

type Handler struct {
	Key interface{}
	Fn  func(data interface{})
}

func NewHandler() *Handler {
	return &Handler{}
}

func CreateHandler(key interface{}, fn func(data interface{})) *Handler {
	return &Handler{
		Key: key,
		Fn:  fn,
	}
}
