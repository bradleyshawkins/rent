package http

import (
	"net/http"
)

func (r *Router) Health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
