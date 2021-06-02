package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/ivchip/go-meli-filter-ip/domain"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	limitCmd, _ := strconv.Atoi(os.Getenv("LIMIT_COMMAND"))
	limitReq, _ := strconv.Atoi(os.Getenv("LIMIT_REQUEST"))
	data := getCountries()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(httprate.LimitByIP(limitReq, time.Duration(limitCmd)*time.Minute))
	c1 := make(chan domain.ResponseIp)
	c2 := make(chan float64)
	// Ip test
	router.Get("/getIp", func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		go getAPIIpFy(ip, c1)
		responseIp := <-c1
		var result domain.ContextualResult
		result.Location = responseIp.Location
		coin := "USD_" + data.Results[responseIp.Location.Country].CurrencyId
		go getAPICurrency(coin, c2)
		cur := <-c2
		result.CurrencyQuote = cur
		data, _ := json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
	http.ListenAndServe(":"+port, router)
}

func getAPIIpFy(ip string, ch chan<- domain.ResponseIp) {
	var responseIp domain.ResponseIp
	urlBase := os.Getenv("API_IPFY")
	url := fmt.Sprintf(urlBase, ip)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		ch <- responseIp
	}
	err = json.NewDecoder(response.Body).Decode(&responseIp)
	if err != nil {
		log.Fatal(err)
		ch <- responseIp
	}
	ch <- responseIp
}

func getAPICurrency(coin string, ch chan<- float64) {
	var currency float64
	result := make(map[string]float64)
	urlAPI := os.Getenv("API_CURRCONV")
	url := fmt.Sprintf(urlAPI, coin)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		ch <- currency
	}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
		ch <- currency
	}
	currency = result[coin]
	ch <- currency
}

func getCountries() domain.Data {
	file, _ := ioutil.ReadFile("countries.json")
	var data domain.Data
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
