package dtos

type HistoricConsumptionDTO struct {
	Client_Id string       `json:"client_id"`
	Measures  []MeasureDTO `json:"measures"`
}

type MeasureDTO struct {
	Consumption uint16 `json:"consumption"`
	Begin       string `json:"begin"`
	End         string `json:"end"`
}
