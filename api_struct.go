package main

type Ticker struct {
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	Symbol             string  `json:"symbol"`
	Rank               int32   `json:"rank,string"`
	Price_usd          float64 `json:"price_usd,string"`
	Price_btc          float64 `json:"price_btc,string"`
	H24h_volume_usd    string  `json:"24h_volume_usd"`
	Market_cap_usd     float64 `json:"market_cap_usd,string"`
	Available_supply   string  `json:"available_supply"`
	Total_supply       string  `json:"total_supply"`
	Max_supply         string  `json:"max_supply"`
	Percent_change_1h  float64 `json:"percent_change_1h,string"`
	Percent_change_24h float64 `json:"percent_change_24h,string"`
	Percent_change_7d  float64 `json:"percent_change_7d,string"`
	Last_updated       int64   `json:"last_updated,string"`
}
