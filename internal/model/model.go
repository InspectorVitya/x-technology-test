package model

type Stocks struct {
	Symbol         string  `json:"symbol" db:"symbol"`
	Price          float64 `json:"price_24h" db:"price"`
	Volume         float64 `json:"volume_24h" db:"volume"`
	LastTradePrice float64 `json:"last_trade_price" db:"last_trade"`
}

type ResponseStocks struct {
	Price     float64 `json:"price"`
	Volume    float64 `json:"volume"`
	LastTrade float64 `json:"last_trade"`
}

type Response map[string]ResponseStocks
