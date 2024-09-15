package estimation

import (
	"energy_estimation/domain/historic_consumption"
	"energy_estimation/domain/tariff"
	"fmt"
	"time"
)

type IEstimationService interface {
	GetDates(historic *historic_consumption.HistoricConsumption) (time.Time, time.Time, error)
	Estimate(historic *historic_consumption.HistoricConsumption, tariffsRules *[]tariff.TariffRule) uint16
}
type EstimationService struct {
	IEstimationService
}

func (s *EstimationService) GetDates(historic *historic_consumption.HistoricConsumption) (time.Time, time.Time, error) {
	if len(historic.Measures) == 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("no measures available")
	}
	dateBegin := time.Now()
	dateEnd := time.Now()
	layout := "2006/01/02"
	for _, measure := range historic.Measures {
		begin, errBegin := time.Parse(layout, measure.Begin)
		end, errEnd := time.Parse(layout, measure.End)
		if errBegin != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("error %v", errBegin)
		}
		if errEnd != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("error %v", errEnd)
		}
		if begin.Before(dateBegin) {
			dateBegin = begin
		}
		if end.After(dateEnd) {
			dateEnd = end
		}
	}

	return dateBegin, dateEnd, nil

}

func (s *EstimationService) Estimate(historic *historic_consumption.HistoricConsumption, tariffsRules *[]tariff.TariffRule) []ConsumptionEstimation {
	totalMeasures := calculateTotalMeasure(historic)
	return applyTariffs(totalMeasures, tariffsRules)

}

func calculateTotalMeasure(historic *historic_consumption.HistoricConsumption) uint16 {
	var totalMeasures uint16 = 0
	for _, measure := range historic.Measures {
		totalMeasures += measure.Consumption

	}
	return totalMeasures
}

func applyTariffs(totalMeasures uint16, tariffsRules *[]tariff.TariffRule) []ConsumptionEstimation {
	estimations := make([]ConsumptionEstimation, len(*tariffsRules))

	for i, estimation := range *tariffsRules {
		estimations[i] = ConsumptionEstimation{
			Id:         estimation.Id,
			Estimation: uint16(float32(totalMeasures) * estimation.Ratio)}
	}
	return estimations
}
