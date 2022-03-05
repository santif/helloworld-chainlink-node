package main

type Request struct {
	ID          string                 `json:"id,omitempty"`
	Data        map[string]interface{} `json:"data"`
	ResponseURL string                 `json:"responseURL,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
}

type PriceResponseData struct {
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}

type PriceResponse struct {
	Data  PriceResponseData `json:"data"`
	Error string            `json:"error,omitempty"`
}
