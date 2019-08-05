package data_resources

import (
	"github.com/noah-blockckain/noh-explorer-api/helpers"
	"github.com/noah-blockckain/noh-explorer-api/resource"
	"github.com/noah-blockckain/noh-explorer-tools/models"
)

type Unbond struct {
	PubKey string `json:"pub_key"`
	Coin   string `json:"coin"`
	Value  string `json:"value"`
}

func (Unbond) Transform(txData resource.ItemInterface, params ...interface{}) resource.Interface {
	data := txData.(*models.UnbondTxData)

	return Unbond{
		PubKey: data.PubKey,
		Coin:   data.Coin,
		Value:  helpers.PipStr2Bip(data.Value),
	}
}
