package apm

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// get service metrics
func (a *Micro) Metrics(w http.ResponseWriter, req *http.Request) {
	promhttp.Handler().ServeHTTP(w, req)
}
