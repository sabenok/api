package data_resources

import (
	"github.com/noah-blockchain/noah-explorer-api/helpers"
	"github.com/noah-blockchain/noah-explorer-api/resource"
	"github.com/noah-blockchain/noah-explorer-tools/models"
)

type CreateCoin struct {
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	InitialAmount        string `json:"initial_amount"`
	InitialReserve       string `json:"initial_reserve"`
	ConstantReserveRatio string `json:"constant_reserve_ratio"`
}

func (CreateCoin) Transform(txData resource.ItemInterface, params ...interface{}) resource.Interface {
	data := txData.(*models.CreateCoinTxData)

	return CreateCoin{
		Name:                 data.Name,
		Symbol:               data.Symbol,
		InitialAmount:        helpers.QNoahStr2Noah(data.InitialAmount),
		InitialReserve:       helpers.QNoahStr2Noah(data.InitialReserve),
		ConstantReserveRatio: data.ConstantReserveRatio,
	}
}
