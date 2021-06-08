package http

import (
	"encoding/json"
	"github.com/ivchip/go-meli-filter-ip/usecases"
	"net"
	"net/http"
)

func GetByIP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	result, err := usecases.GetByIP(ip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
