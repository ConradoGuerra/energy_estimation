package historic_consomation

type HistoricConsomation struct {
	Client_Id string
	Measures  []Measure
}

type Measure struct {
	Consomation uint16
	Begin       string
	End         string
}
