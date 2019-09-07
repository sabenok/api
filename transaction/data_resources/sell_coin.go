package data_resources

import (
	"github.com/noah-blockchain/noah-explorer-api/helpers"
	"github.com/noah-blockchain/noah-explorer-api/resource"
	"github.com/noah-blockchain/noah-explorer-tools/models"
)

type SellCoin struct {
	CoinToSell        string `json:"coin_to_sell"`
	CoinToBuy         string `json:"coin_to_buy"`
	ValueToSell       string `json:"value_to_sell"`
	ValueToBuy        string `json:"value_to_buy"`
	MinimumValueToBuy string `json:"minimum_value_to_buy"`
}

func (SellCoin) Transform(txData resource.ItemInterface, params ...interface{}) resource.Interface {
	data := txData.(*models.SellCoinTxData)
	model := params[0].(models.Transaction)

	return SellCoin{
		CoinToSell:        data.CoinToSell,
		CoinToBuy:         data.CoinToBuy,
		ValueToSell:       helpers.QNoahStr2Noah(data.ValueToSell),
		ValueToBuy:        helpers.QNoahStr2Noah(model.Tags["tx.return"]),
		MinimumValueToBuy: helpers.QNoahStr2Noah(data.MinimumValueToBuy),
	}
}
