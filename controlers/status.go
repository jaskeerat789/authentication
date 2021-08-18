package controlers

import (
	"net/http"

	"github.com/hashicorp/go-hclog"
)

type StatusController struct {
	l hclog.Logger
}

func NewStatusCotroller() *StatusController {
	log := hclog.New(&hclog.LoggerOptions{
		Name: "Status controller",
	})
	return &StatusController{l: log}
}

func (s StatusController) GetStatus(w http.ResponseWriter, r *http.Request) {
	s.l.Info("GET /status")
	w.Write([]byte("API is running"))
}
