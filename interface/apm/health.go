package apm

import (
	"net/http"
)

// check service health
func (a *Micro) Health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("healthy"))
}
