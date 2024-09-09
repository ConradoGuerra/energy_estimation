package services

import (
	"energy_estimation/src/domain/estimation"
	"energy_estimation/src/domain/historic_consomation"
	"energy_estimation/src/domain/tariff"
	"fmt"
)

func NewEstimationService(historic *historic_consomation.HistoricConsomation) (estimation.Estimation, error) {
	service := &estimation.EstimationService{}
	begin, end, error := service.GetDates(historic)
	if error != nil {
		return estimation.Estimation{}, fmt.Errorf("error %v", error)
	}
	estimations := service.Estimate(historic, tariff.TariffsRules)

	return estimation.Estimation{
		Begin:                  begin,
		End:                    end,
		ConsomationsEstimation: estimations}, nil
}
