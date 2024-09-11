package estimation

import "time"

type Estimation struct {
	Begin                  time.Time
	End                    time.Time
	ConsumptionEstimations []ConsumptionEstimation
}

type ConsumptionEstimation struct {
	Id         string
	Estimation uint16
}
