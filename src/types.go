package main

type BinanceRequestBody struct {
	Method string    `json:"method"`
	Params [1]string `json:"params"`
	ID     int       `json:"id"`
}

type BinanceResponseBody struct {
	EventType     string `json:"e"`
	EventTime     int    `json:"E"`
	Symbol        string `json:"s"`
	AggTradeID    int    `json:"a"`
	Price         string `json:"p"`
	Quantity      string `json:"q"`
	FirstTradeID  int    `json:"f"`
	LastTradeID   int    `json:"l"`
	TradeTime     int    `json:"T"`
	IsBuyerMMaker bool   `json:"m"`
	Ignore        bool   `json:"M"`
}
