package mlpapi

import (
	"net/http"
	"time"
)

type Rule struct {
	Description  string
	DoesSlotPass func(Slot) bool
}

type MLPClient struct {
	httpClient *http.Client
}

type Pitch struct {
	ID   string
	Path string
	City string
	Name string
}

type MLPResponse struct {
	Meta struct {
		TotalItems int `json:"total_items"`
		Filter     struct {
			Starts string `json:"starts"`
			Ends   string `json:"ends"`
		} `json:"filter"`
	} `json:"meta"`
	Data MLPData `json:"data"`
}

type MLPData []Slot

type Slot struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Starts         time.Time `json:"starts"`
		Ends           time.Time `json:"ends"`
		Price          string    `json:"price"`
		AdminFee       string    `json:"admin_fee"`
		Currency       string    `json:"currency"`
		Availabilities int       `json:"availabilities"`
	} `json:"attributes"`
}
