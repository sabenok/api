package market

import (
	"github.com/noah-blockchain/noah-explorer-api/noahdev"
	"github.com/noah-blockchain/noah-explorer-api/noahdev/responses"
)

type Service struct {
	api      *noahdev.Api
	baseCoin string
}

func NewService(noahDevApi *noahdev.Api, baseCoin string) *Service {
	return &Service{
		api:      noahDevApi,
		baseCoin: baseCoin,
	}
}

type PriceChange struct {
	Price  float64
	Change float64
}

func (s *Service) GetCurrentFiatPriceChange(coin string, currency string) (*PriceChange, error) {
	if coin == s.baseCoin && currency == USDTicker {
		//response, err := s.api.GetCurrentPrice()
		//if err != nil {
		//	return nil, err
		//}

		response := responses.CurrentPriceResponse{} // TODO fix
		response.Data.Price = 1000
		response.Data.Delta = 1.0

		return &PriceChange{
			Price:  response.Data.Price / 10000,
			Change: response.Data.Delta,
		}, nil
	}

	return nil, nil
}
