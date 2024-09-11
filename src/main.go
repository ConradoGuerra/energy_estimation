package main

import (
	"encoding/json"
	"energy_estimation/src/domain/historic_consumption"
	"energy_estimation/src/services"
	"fmt"
	"io"
	"net/http"
)

type HistoricConsumptionDTO struct {
	Client_Id string       `json:"client_id"`
	Measures  []MeasureDTO `json:"measures"`
}

type MeasureDTO struct {
	Consumption uint16 `json:"consumption"`
	Begin       string `json:"begin"`
	End         string `json:"end"`
}

func EstimationHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	var historicConsumptionDTO HistoricConsumptionDTO
	json.Unmarshal(body, &historicConsumptionDTO)

	historicConsumption := convertToDomain(historicConsumptionDTO)

	estimation, err := services.NewEstimationService(&historicConsumption)

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

func convertToDomain(dto HistoricConsumptionDTO) historic_consumption.HistoricConsumption {
	measures := make([]historic_consumption.Measure, len(dto.Measures))
	for i, measure := range dto.Measures {
		measures[i] = historic_consumption.Measure{Begin: measure.Begin, End: measure.End, Consumption: measure.Consumption}
	}
	return historic_consumption.HistoricConsumption{
		Client_Id: dto.Client_Id,
		Measures:  measures,
	}

}

func main() {

	http.HandleFunc("/api/estimation", EstimationHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)

}
