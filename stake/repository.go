package stake

import (
	"github.com/go-pg/pg"
	"github.com/noah-blockchain/noah-explorer-api/helpers"
	"github.com/noah-blockchain/noah-explorer-api/tools"
	"github.com/noah-blockchain/noah-explorer-tools/models"
)

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Get paginated list of stakes by Noah address
func (repository Repository) GetByAddress(address string, pagination *tools.Pagination) []models.Stake {
	var stakes []models.Stake
	var err error

	pagination.Total, err = repository.db.Model(&stakes).
		Column("Coin.symbol", "Validator.public_key", "OwnerAddress._").
		Where("owner_address.address = ?", address).
		Apply(pagination.Filter).
		SelectAndCount()

	helpers.CheckErr(err)

	return stakes
}

// Get total delegated noah value
func (repository Repository) GetSumInNoahValue() (string, error) {
	var sum string
	err := repository.db.Model(&models.Stake{}).ColumnExpr("SUM(noah_value)").Select(&sum)
	return sum, err
}

// Get total delegated sum by address
func (repository Repository) GetSumInNoahValueByAddress(address string) (string, error) {
	var sum string
	err := repository.db.Model(&models.Stake{}).
		Column("OwnerAddress._").
		ColumnExpr("SUM(noah_value)").
		Where("owner_address.address = ?", address).
		Select(&sum)

	return sum, err
}
