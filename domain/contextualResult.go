package domain

type location struct {
	Country    string  `json:"country"`
	Region     string  `json:"region"`
	City       string  `json:"city"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	PostalCode string  `json:"postalCode"`
	Timezone   string  `json:"timezone"`
	GeoNameId  int     `json:"geonameId"`
}

type as struct {
	Asn    int    `json:"asn"`
	Name   string `json:"name"`
	Route  string `json:"route"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

type proxy struct {
	Proxy bool `json:"proxy"`
	Vpn   bool `json:"vpn"`
	Tor   bool `json:"tor"`
}

type ResponseIp struct {
	Ip       string   `json:"ip"`
	Location location `json:"location"`
	Domains  []string `json:"domains"`
	As       as       `json:"as"`
	Isp      string   `json:"isp"`
	Proxy    proxy    `json:"proxy"`
}

type countries struct {
	Alpha3         string `json:"alpha3"`
	CurrencyId     string `json:"currencyId"`
	CurrencyName   string `json:"currencyName"`
	CurrencySymbol string `json:"currencySymbol"`
	Id             string `json:"id"`
	Name           string `json:"name"`
}

type Data struct {
	Results map[string]countries `json:"results"`
	Note    string               `json:"note"`
}

type ContextualResult struct {
	Location      location `json:"location"`
	CurrencyQuote float64  `json:"currencyQuote"`
}

type ContextualResultUseCases interface {
	GetByIP(ip string) (ContextualResult, error)
}
