package usecases

import (
	"encoding/json"
	"fmt"
	"github.com/ivchip/go-meli-filter-ip/domain"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetByIP(ip string) (domain.ContextualResult, error) {
	var result domain.ContextualResult
	countries := getCountries()
	responseIp, err := getIp(ip)
	if err != nil {
		return result, err
	}
	result.Location = responseIp.Location
	coin := "USD_" + countries.Results[responseIp.Location.Country].CurrencyId
	currency, err := getCurrency(coin)
	if err != nil {
		return result, err
	}
	result.CurrencyQuote = currency
	return result, nil
}

func getIp(ip string) (r domain.ResponseIp, err error) {
	resultOut, resultErr := getAPIIpFy(ip)
	var open bool
	if err, open = <-resultErr; open {
		return
	}
	r = <-resultOut
	return
}

func getAPIIpFy(ip string) (<-chan domain.ResponseIp, <-chan error) {
	var responseIp domain.ResponseIp
	urlBase := os.Getenv("API_IPFY")
	url := fmt.Sprintf(urlBase, ip)
	out := make(chan domain.ResponseIp, 1)
	errs := make(chan error, 1)
	go func() {
		response, err := http.Get(url)
		if err != nil {
			errs <- err
		}
		err = json.NewDecoder(response.Body).Decode(&responseIp)
		if err != nil {
			errs <- err
		}
		out <- responseIp
		close(out)
		close(errs)
	}()
	return out, errs
}

func getCurrency(coin string) (currency float64, err error) {
	resultOut, resultErr := getAPICurrency(coin)
	var open bool
	if err, open = <-resultErr; open {
		return
	}
	currency = <-resultOut
	return
}

func getAPICurrency(coin string) (<-chan float64, <-chan error) {
	var currency float64
	result := make(map[string]float64)
	urlAPI := os.Getenv("API_CURRCONV")
	url := fmt.Sprintf(urlAPI, coin)
	out := make(chan float64, 1)
	errs := make(chan error, 1)
	go func() {
		response, err := http.Get(url)
		if err != nil {
			errs <- err
		}
		err = json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			errs <- err
		}
		currency = result[coin]
		out <- currency
		close(out)
		close(errs)
	}()
	return out, errs
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
