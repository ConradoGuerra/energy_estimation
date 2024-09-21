package tariff

type TariffRepository interface {
	GetTariffs() *[]TariffRule
}

type TariffService struct {
	TariffRepository TariffRepository
}

func NewService(tariffRepository TariffRepository) *TariffService {
	return &TariffService{tariffRepository}
}

func (s *TariffService) GetTariffs() *[]TariffRule {
	return s.TariffRepository.GetTariffs()
}
