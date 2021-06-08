package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis/v8"
	"github.com/ivchip/go-meli-filter-ip/domain"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	redisLimiter "github.com/ulule/limiter/v3/drivers/store/redis"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	redisAddr := os.Getenv("REDIS_ADDRESS")
	redisPass := os.Getenv("REDIS_PASS")
	limitCmd := os.Getenv("LIMIT_COMMAND")
	data := getCountries()

	// Define a limit rate to number requests per hour.
	rate, err := limiter.NewRateFromFormatted(limitCmd)
	if err != nil {
		log.Fatal(err)
		return
	}

	client := redis.NewClient(&redis.Options{Addr: redisAddr, Password: redisPass, DB: 0})

	// Create a store with the redis client.
	store, err := redisLimiter.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix: "limiter_chi",
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a new middleware with the limiter instance.
	middlewareLimit := stdlib.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)))

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middlewareLimit.Handler)

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
