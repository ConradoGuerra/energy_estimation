package dtos

type HistoricConsumption struct {
	Client_Id string
	Measures  []Measure
}

type Measure struct {
	Consumption uint16
	Begin       string
	End         string
}
