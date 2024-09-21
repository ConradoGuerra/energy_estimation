package in_memory

import "energy_estimation/domain/tariff"

type InMemoryTariffRepo struct{}

func (i InMemoryTariffRepo) GetTariffs() *[]tariff.TariffRule {
	return &[]tariff.TariffRule{
		{Id: "BASE", Ratio: 1},
		{Id: "OFF-PEAK", Ratio: .5},
		{Id: "PEAK", Ratio: .5},
		{Id: "CUSTOM", Ratio: .67},
	}
}
