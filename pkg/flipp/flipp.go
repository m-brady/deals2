package flipp

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const locale = "en-ca"
const postalCode = "M4P1V6"
const host = "https://gateflipp.flippback.com/bf/flipp"
const merchants = host + "/merchants"
const flyers = host + "/flyers"

type Flyer struct {
	Id            int    `json:"id"`
	MerchantId    int    `json:"merchant_id"`
	ValidTo       string `json:"valid_to"`
	ValidFrom     string `json:"valid_from"`
	AvailableTo   string `json:"available_to"`
	AvailableFrom string `json:"available_from"`
}

type Response struct {
	RefreshedAt string     `json:"refreshed_at"`
	Merchants   []Merchant `json:"merchants"`
	Flyers      []Flyer    `json:"flyers"`
	Items       []Item     `json:"items"`
}

type Merchant struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	UsBased        bool   `json:"us_based"`
	NameIdentifier string `json:"name_identifier"`
}

type Item struct {
	Id           int64   `json:"id"`
	FlyerItemId  int64   `json:"flyer_item_id"`
	FlyerId      int     `json:"flyer_id"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
	ValidTo      string  `json:"valid_to"`
	ValidFrom    string  `json:"valid_from"`
	MerchantId   int     `json:"merchant_id"`
	Price        float64 `json:"price,string"`
}

func get(url string) Response {
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()

	q.Add("locale", locale)
	q.Add("postal_code", postalCode)
	req.URL.RawQuery = q.Encode()

	log.Printf("Performing request: %#v", req)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response Response
	_ = decoder.Decode(&response)
	log.Println(response)
	return response
}

func GetMerchants() Response {
	return get(merchants)
}

func GetFlyers() Response {
	return get(flyers)
}

func GetFlyer(flyerId int64) Response {
	return get(flyers + "/" + strconv.FormatInt(flyerId, 10))
}
