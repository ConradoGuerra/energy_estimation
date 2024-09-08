package services

import (
	"energy_estimation/src/domain/estimation"
	"energy_estimation/src/domain/historic_consomation"
	"fmt"
)

func NewEstimationService(historic *historic_consomation.HistoricConsomation) (estimation.Estimation, error) {
	service := &estimation.EstimationService{}
	begin, end, error := service.GetDates(historic)
	if error != nil {
		return estimation.Estimation{}, fmt.Errorf("error %v", error)
	}
	totalEstimation := service.Estimate(historic)

	return estimation.Estimation{
		Pdl:          historic.Client_Id,
		EstimationId: "estimationId",
		Begin:        begin,
		End:          end,
		Estimation:   totalEstimation}, nil
}
