package flipp

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
