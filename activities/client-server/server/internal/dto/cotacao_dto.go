package dto

import "time"

type CotacaoDTO struct {
	code       string    `json:"code"`
	codeIn     string    `json:"codein"`
	name       string    `json:"name"`
	high       float64   `json:"high"`
	low        float64   `json:"low"`
	varBid     float64   `json:"varBid"`
	pctChange  float64   `json:"pctChange"`
	bid        float64   `json:"bid"`
	ask        float64   `json:"ask"`
	timestamp  time.Time `json:"timestamp"`
	createDate time.Time `json:"createDate"`
}

func (cotacao *CotacaoDTO) GetCode() string {
	return cotacao.code
}

func (cotacao *CotacaoDTO) GetCodeIn() string {
	return cotacao.codeIn
}

func (cotacao *CotacaoDTO) GetName() string {
	return cotacao.name
}

func (cotacao *CotacaoDTO) GetHigh() float64 {
	return cotacao.high
}

func (cotacao *CotacaoDTO) GetLow() float64 {
	return cotacao.low
}

func (cotacao *CotacaoDTO) GetVarBid() float64 {
	return cotacao.varBid
}

func (cotacao *CotacaoDTO) GetPctChange() float64 {
	return cotacao.pctChange
}

func (cotacao *CotacaoDTO) GetBid() float64 {
	return cotacao.bid
}

func (cotacao *CotacaoDTO) GetAsk() float64 {
	return cotacao.ask
}
