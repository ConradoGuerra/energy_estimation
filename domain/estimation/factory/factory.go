package factory

import (
	"energy_estimation/domain/estimation"
	"energy_estimation/domain/historic_consumption"
	"energy_estimation/domain/tariff"
	"fmt"
)

func CreateEstimationService(historic *historic_consumption.HistoricConsumption) (estimation.Estimation, error) {
	service := &estimation.EstimationService{}
	begin, end, error := service.GetDates(historic)
	if error != nil {
		return estimation.Estimation{}, fmt.Errorf("error %v", error)
	}
	estimations := service.Estimate(historic, tariff.TariffsRules)

	return estimation.Estimation{
		Begin:                  begin,
		End:                    end,
		ConsumptionEstimations: estimations}, nil
}
