package estimation

import "time"

type Estimation struct {
	Begin                  time.Time
	End                    time.Time
	ConsomationsEstimation []ConsomationEstimation
}

type ConsomationEstimation struct {
	Id         string
	Estimation uint16
}
