package http

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/ivchip/go-meli-filter-ip/domain"
	"github.com/ivchip/go-meli-filter-ip/usecases"
	"net"
	"net/http"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

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

func GetByIp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ip := chi.URLParam(r, "ip")
	result, err := usecases.GetByIp(ip)
	if err != nil {
		response := newResponse(Error, "The Ip is empty", nil)
		data, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}
	response := newResponse(Message, "Success", result)
	data, _ := json.Marshal(&response)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := usecases.GetAll()
	if err != nil {
		response := newResponse(Error, "The Ip is empty", nil)
		data, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}
	response := newResponse(Message, "Success", result)
	data, _ := json.Marshal(&response)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body domain.IpBlocking
	json.NewDecoder(r.Body).Decode(&body)
	err := usecases.Create(&body)
	if err != nil {
		response := newResponse(Error, err.Error(), nil)
		data, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}
	response := newResponse(Message, "Success", nil)
	data, _ := json.Marshal(&response)
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ip := chi.URLParam(r, "ip")
	var body domain.IpBlocking
	json.NewDecoder(r.Body).Decode(&body)
	err := usecases.Update(ip, &body)
	if err != nil {
		response := newResponse(Error, err.Error(), nil)
		data, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}
	response := newResponse(Message, "Success", nil)
	data, _ := json.Marshal(&response)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
