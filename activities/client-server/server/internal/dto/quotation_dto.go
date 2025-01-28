package dto

import "time"

type QuotationDTO struct {
	CurrencyDTO CurrencyDTO `json:"USDBRL"`
}

type CurrencyDTO struct {
	Code       string    `json:"code"`
	CodeIn     string    `json:"codein"`
	Name       string    `json:"name"`
	High       string    `json:"high"`
	Low        string    `json:"low"`
	VarBid     string    `json:"varBid"`
	PctChange  string    `json:"pctChange"`
	Bid        string    `json:"bid"`
	Ask        string    `json:"ask"`
	Timestamp  string    `json:"timestamp"`
	CreateDate time.Time `json:"create_date"`
}
