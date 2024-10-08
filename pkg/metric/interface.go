package metric

import "time"

//HTTP application
type HTTP struct {
	Handler    string
	Method     string
	StatusCode string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

//NewHTTP create a new HTTP app
func NewHTTP(handler string, method string) *HTTP {
	return &HTTP{
		Handler: handler,
		Method:  method,
	}
}

//Started start monitoring the app
func (h *HTTP) Started() {
	h.StartedAt = time.Now()
}

// Finished app finished
func (h *HTTP) Finished() {
	h.FinishedAt = time.Now()
	h.Duration = time.Since(h.StartedAt).Seconds()
}

//IService definition
type IService interface {
	SaveHTTP(h *HTTP)
}
