package response

import "server/domain/Entity"

type CotacaoResponse struct {
	UsdParaBrl Entity.Cotacao `json:"USDBRL"`
}

type CotacaoBidResponse struct {
	Bid string `json:"bid"`
}
