package servicemocks

import "auth-and-db-service/services"

type MockSearchService struct{}

func (mss *MockSearchService) AddProductToSearchService(p services.SearchProduct) error {

	return nil
}
