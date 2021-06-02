package main

import (
	"encoding/json"
	"fmt"
	limiter "github.com/davidleitw/gin-limiter"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/heroku/x/hmetrics/onload"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type location struct {
	Country string `json:"country"`
	Region string `json:"region"`
	City string `json:"city"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	PostalCode string `json:"postalCode"`
	Timezone string `json:"timezone"`
	GeoNameId int `json:"geonameId"`
}

type as struct {
	Asn int `json:"asn"`
	Name string `json:"name"`
	Route string `json:"route"`
	Domain string `json:"domain"`
	Type string `json:"type"`
}

type proxy struct {
	Proxy bool `json:"proxy"`
	Vpn bool `json:"vpn"`
	Tor bool `json:"tor"`
}

type ResponseIp struct {
	Ip string `json:"ip"`
	Location location `json:"location"`
	Domains []string `json:"domains"`
	As as `json:"as"`
	Isp string `json:"isp"`
	Proxy proxy `json:"proxy"`
}

type countries struct {
	Alpha3 string `json:"alpha3"`
	CurrencyId string `json:"currencyId"`
	CurrencyName string `json:"currencyName"`
	CurrencySymbol string `json:"currencySymbol"`
	Id string `json:"id"`
	Name string `json:"name"`
}

type Data struct {
	Results map[string]countries `json:"results"`
	Note string `json:"note"`
}

type ContextualResult struct {
	Location location `json:"location"`
	CurrencyQuote float64 `json:"currencyQuote"`
}

func main() {
	port := os.Getenv("SERVER_PORT")
	redisAddr := os.Getenv("REDIS_ADDRESS")
	redisPass := os.Getenv("REDIS_PASS")
	limitCmd := os.Getenv("LIMIT_COMMAND")
	limitReq, _ := strconv.Atoi(os.Getenv("LIMIT_REQUEST"))
	data := getCountries()
	// Set redis client
	rdb := redis.NewClient(&redis.Options{Addr: redisAddr, Password: redisPass, DB: 0})
	// Set limit middleWare
	dispatcher, err := limiter.LimitDispatcher(limitCmd, limitReq, rdb)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.New()
	router.Use(gin.Logger())
	c1 := make(chan ResponseIp)
	c2 := make(chan float64)
	// Ip test
	router.GET("/getIp", dispatcher.MiddleWare(limitCmd, limitReq), func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		go getAPIIpFy(ip, c1)
		responseIp := <- c1
		var result ContextualResult
		result.Location = responseIp.Location
		coin := "USD_" + data.Results[responseIp.Location.Country].CurrencyId
		go getAPICurrency(coin, c2)
		cur := <- c2
		result.CurrencyQuote = cur
		ctx.JSON(http.StatusOK, result)
	})

	router.Run(":" + port)
}

func getAPIIpFy(ip string, ch chan <- ResponseIp) {
	var responseIp ResponseIp
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

func getAPICurrency(coin string, ch chan <- float64)  {
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

func getCountries() Data {
	file, _ := ioutil.ReadFile("countries.json")
	var data Data
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
