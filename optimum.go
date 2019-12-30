package deals

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

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

type Flyer struct {
	Id            int    `json:"id"`
	MerchantId    int    `json:"merchant_id"`
	ValidTo       string `json:"valid_to"`
	ValidFrom     string `json:"valid_from"`
	AvailableTo   string `json:"available_to"`
	AvailableFrom string `json:"available_from"`
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

const locale = "en-ca"
const postalCode = "M4P1V6"
const host = "https://gateflipp.flippback.com/bf/flipp"
const merchants = host + "/merchants"
const flyers = host + "/flyers"

func main() {
	//makeRequest("loblaws")
	//mb()
	//loadMerchants()
	//loadFlyers()
	loadFlyerItems(3184499)
}

func makeRequest(query string) Response {
	req, _ := http.NewRequest("GET", host, nil)
	q := req.URL.Query()

	q.Add("locale", locale)
	q.Add("postal_code", postalCode)
	q.Add("q", query)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL)

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp)

	decoder := json.NewDecoder(resp.Body)

	var response Response

	_ = decoder.Decode(&response)
	fmt.Println(response)
	return response
}

func getMerchants() Response {
	req, _ := http.NewRequest("GET", merchants, nil)
	fmt.Println(req.URL)

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp)

	decoder := json.NewDecoder(resp.Body)

	var response Response

	_ = decoder.Decode(&response)
	fmt.Println(response)
	return response
}

func getFlyers() Response {
	req, _ := http.NewRequest("GET", flyers, nil)
	q := req.URL.Query()

	q.Add("locale", locale)
	q.Add("postal_code", postalCode)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL)

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp)

	decoder := json.NewDecoder(resp.Body)

	var response Response

	_ = decoder.Decode(&response)
	fmt.Println(response)
	return response
}

func getFlyer(flyerId int64) Response {
	fmt.Println(flyerId)
	req, _ := http.NewRequest("GET", flyers+"/"+strconv.FormatInt(flyerId, 10), nil)
	fmt.Println(req.URL)

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp)

	decoder := json.NewDecoder(resp.Body)

	var response Response

	_ = decoder.Decode(&response)
	fmt.Println(response)
	return response
}

func mb() {

	db, err := sql.Open("mysql", "root:brady@/optimum")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}

func insertMerchants(merchants []Merchant) {

	db, err := sql.Open("mysql", "root:brady@/optimum")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:

	for _, merchant := range merchants {
		db.Exec("insert ignore into merchant (id, us_based, name_identifier, name) values (?, ?, ?, ?)",
			merchant.Id, merchant.UsBased, merchant.NameIdentifier, merchant.Name)
	}
}

func insertFlyers(flyers []Flyer) {

	db, err := sql.Open("mysql", "root:brady@/optimum")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:

	for _, flyer := range flyers {
		db.Exec("insert ignore into flyer (id, merchant_id, valid_from, valid_to, available_from, available_to) values (?, ?, ?, ?, ? ,?)",
			flyer.Id, flyer.MerchantId, flyer.ValidFrom, flyer.ValidTo, flyer.AvailableFrom, flyer.AvailableTo)
	}
}

func insertFlyerItems(items []Item) {

	db, err := sql.Open("mysql", "root:brady@/optimum")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:

	for _, item := range items {
		fmt.Println(item)
		//price, _ := strconv.ParseFloat(item.Price, 64)
		db.Exec("insert ignore into flyer_item (flyer_item_id, name, flyer_id, current_price, valid_to, valid_from) values (?, ?, ?, ?, ? ,?)",
			item.Id, item.Name, item.FlyerId, item.Price, item.ValidTo, item.ValidFrom)
	}
}

func loadMerchants() {
	response := getMerchants()
	insertMerchants(response.Merchants)
}

func loadFlyers() {
	response := getFlyers()
	insertFlyers(response.Flyers)
}

func loadFlyerItems(flyerId int64) {
	response := getFlyer(flyerId)
	insertFlyerItems(response.Items)
}





