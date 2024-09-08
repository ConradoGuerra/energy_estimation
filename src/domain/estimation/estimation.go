package estimation

import "time"

type Estimation struct {
	Pdl          string
	EstimationId string
	Begin        time.Time
	End          time.Time
	Estimation   uint16
}
