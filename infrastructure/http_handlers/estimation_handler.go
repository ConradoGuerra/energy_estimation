package http_handlers

import (
	"encoding/json"
	"energy_estimation/domain/estimation/factory"
	"energy_estimation/domain/historic_consumption"
	"energy_estimation/infrastructure/dtos"
	"io"
	"net/http"
)

func EstimationHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	var historicConsumptionDTO dtos.HistoricConsumptionDTO
	json.Unmarshal(body, &historicConsumptionDTO)

	historicConsumption := convertToDomain(historicConsumptionDTO)

	estimation, err := factory.CreateEstimationService(&historicConsumption)

	if err != nil {
		http.Error(w, "Error calculating estimation", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	jsonResponse, err := json.Marshal(estimation)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)

}

func convertToDomain(dto dtos.HistoricConsumptionDTO) historic_consumption.HistoricConsumption {
	measures := make([]historic_consumption.Measure, len(dto.Measures))
	for i, measure := range dto.Measures {
		measures[i] = historic_consumption.Measure{Begin: measure.Begin, End: measure.End, Consumption: measure.Consumption}
	}
	return historic_consumption.HistoricConsumption{
		Client_Id: dto.Client_Id,
		Measures:  measures,
	}

}
