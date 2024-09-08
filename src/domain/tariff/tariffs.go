package tariff

type TariffRule struct {
	Id    string
	Ratio float32
}

var TariffsRules = &[]TariffRule{
	{Id: "BASE", Ratio: 1},
	{Id: "OFF-PEAK", Ratio: .5},
	{Id: "PEAK", Ratio: .5},
	{Id: "CUSTOM", Ratio: .67},
}
