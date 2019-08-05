package delegation

import (
	"github.com/noah-blockchain/noah-explorer-api/helpers"
	"github.com/noah-blockchain/noah-explorer-api/resource"
	"github.com/noah-blockchain/noah-explorer-tools/models"
)

type Resource struct {
	Coin     string `json:"coin"`
	PubKey   string `json:"pub_key"`
	Value    string `json:"value"`
	BipValue string `json:"bip_value"`
}

func (resource Resource) Transform(model resource.ItemInterface, params ...interface{}) resource.Interface {
	stake := model.(models.Stake)

	return Resource{
		Coin:     stake.Coin.Symbol,
		PubKey:   stake.Validator.GetPublicKey(),
		Value:    helpers.PipStr2Bip(stake.Value),
		BipValue: helpers.PipStr2Bip(stake.BipValue),
	}
}
