package main

import (
	"energy_estimation/src/domain/historic_consumption"
	"energy_estimation/src/services"
	"fmt"
)

func main() {

	var historic = historic_consumption.HistoricConsumption{
		Client_Id: "historic_pdl",
		Measures: []historic_consumption.Measure{
			{Consumption: 23, Begin: "2024/09/01", End: "2024/09/30"},
			{Consumption: 4, Begin: "2024/08/01", End: "2024/08/31"},
			{Consumption: 54, Begin: "2024/07/01", End: "2024/07/31"}},
	}

	estimation, _ := services.NewEstimationService(&historic)
	fmt.Print(estimation)

}
