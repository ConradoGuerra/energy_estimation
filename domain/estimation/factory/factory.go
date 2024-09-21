package factory

import (
	"energy_estimation/domain/estimation"
	"energy_estimation/domain/historic_consumption"
	"energy_estimation/infrastructure/repositories/in_memory"
	"fmt"
)

func CreateEstimationService(historic *historic_consumption.HistoricConsumption) (estimation.Estimation, error) {
	service := &estimation.EstimationService{}
	begin, end, error := service.GetDates(historic)
	if error != nil {
		return estimation.Estimation{}, fmt.Errorf("error %v", error)
	}
	tariffRepository := in_memory.InMemoryTariffRepo{}
	estimations := service.Estimate(historic, tariffRepository.GetTariffs())

	return estimation.Estimation{
		Begin:                  begin,
		End:                    end,
		ConsumptionEstimations: estimations}, nil
}
