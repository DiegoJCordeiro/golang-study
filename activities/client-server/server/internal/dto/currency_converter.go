package dto

import (
	"encoding/json"
	"fmt"
	"time"
)

type currencyAux struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type quotationAux struct {
	Currency currencyAux `json:"USDBRL"`
}

var (
	Quotation quotationAux = quotationAux{}
)

func (c *QuotationDTO) UnmarshalJSON(data []byte) error {

	const layout = "2006-01-02 15:04:05"

	if err := json.Unmarshal(data, &Quotation); err != nil {
		return err
	}

	var err error

	c.CurrencyDTO.Code = Quotation.Currency.Code
	c.CurrencyDTO.CodeIn = Quotation.Currency.CodeIn
	c.CurrencyDTO.Name = Quotation.Currency.Name
	c.CurrencyDTO.High = Quotation.Currency.High
	c.CurrencyDTO.Low = Quotation.Currency.Low
	c.CurrencyDTO.VarBid = Quotation.Currency.VarBid
	c.CurrencyDTO.PctChange = Quotation.Currency.PctChange
	c.CurrencyDTO.Bid = Quotation.Currency.Bid
	c.CurrencyDTO.Ask = Quotation.Currency.Ask
	c.CurrencyDTO.Timestamp = Quotation.Currency.Timestamp
	c.CurrencyDTO.CreateDate, err = time.Parse(layout, Quotation.Currency.CreateDate)

	if err != nil {
		return fmt.Errorf("falha ao parsear CreateDate: %v", err)
	}

	return nil
}
