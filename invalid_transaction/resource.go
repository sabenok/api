package invalid_transaction

import (
	"github.com/noah-blockchain/noah-explorer-api/resource"
	"github.com/noah-blockchain/noah-explorer-tools/models"
	"time"
)

type Resource struct {
	Hash      string `json:"hash"`
	Block     uint64 `json:"block"`
	Timestamp string `json:"timestamp"`
	Type      uint8  `json:"type"`
	From      string `json:"from"`
}

func (Resource) Transform(model resource.ItemInterface, params ...interface{}) resource.Interface {
	tx := model.(models.InvalidTransaction)

	return Resource{
		Hash:      tx.GetHash(),
		Block:     tx.BlockID,
		Timestamp: tx.CreatedAt.Format(time.RFC3339),
		Type:      tx.Type,
		From:      tx.FromAddress.GetAddress(),
	}
}
