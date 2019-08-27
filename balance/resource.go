package balance

import (
	"github.com/noah-blockchain/noah-explorer-api/helpers"
	"github.com/noah-blockchain/noah-explorer-api/resource"
	"github.com/noah-blockchain/noah-explorer-tools/models"
)

type Resource struct {
	Coin   string `json:"coin"`
	Amount string `json:"amount"`
}

func (Resource) Transform(model resource.ItemInterface, params ...interface{}) resource.Interface {
	balance := model.(models.Balance)

	return Resource{
		Coin:   balance.Coin.Symbol,
		Amount: helpers.QNoahStr2Noah(balance.Value),
	}
}
