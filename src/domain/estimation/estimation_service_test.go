package estimation_test

import (
	"energy_estimation/src/domain/estimation"
	"energy_estimation/src/domain/historic_consomation"
	"energy_estimation/src/domain/tariff"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEstimationService_GetDates(t *testing.T) {
	var testCases = []struct {
		name          string
		historic      historic_consomation.HistoricConsomation
		expectedBegin time.Time
		expectedEnd   time.Time
		expectedError error
	}{
		{
			name: "should return earliest begin and latest end from measures",
			historic: historic_consomation.HistoricConsomation{
				Client_Id: "Client",
				Measures: []historic_consomation.Measure{
					{
						Consomation: 12,
						Begin:       "2024/08/01",
						End:         "2024/08/31",
					},
					{
						Consomation: 5,
						Begin:       "2024/05/01",
						End:         "2024/05/31",
					},
					{
						Consomation: 67,
						Begin:       "2024/10/01",
						End:         "2024/10/31",
					},
				},
			},
			expectedBegin: time.Date(2024, time.May, 1, 0, 0, 0, 0, time.UTC),
			expectedEnd:   time.Date(2024, time.October, 31, 0, 0, 0, 0, time.UTC),
			expectedError: nil,
		},
		{
			name: "should return an error when no measure is present",
			historic: historic_consomation.HistoricConsomation{
				Client_Id: "Client",
				Measures:  []historic_consomation.Measure{},
			},
			expectedBegin: time.Time{},
			expectedEnd:   time.Time{},
			expectedError: errors.New("no measures available"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert := assert.New(t)

			estimationService := estimation.EstimationService{}
			gotBegin, gotEnd, err := estimationService.GetDates(&testCase.historic)

			if testCase.expectedError != nil {
				assert.EqualError(err, testCase.expectedError.Error())
			} else {
				assert.NoError(err)
			}

			assert.Equal(testCase.expectedBegin, gotBegin)
			assert.Equal(testCase.expectedEnd, gotEnd)
		})
	}
}

func TestEstimationService_Estimate(t *testing.T) {
	testCases := []struct {
		name               string
		historic           historic_consomation.HistoricConsomation
		expectedEstimation []estimation.ConsomationEstimation
	}{
		{
			name: "should return all expected estimations",
			historic: historic_consomation.HistoricConsomation{
				Client_Id: "Client",
				Measures: []historic_consomation.Measure{
					{
						Consomation: 12,
						Begin:       "2024/08/01",
						End:         "2024/08/31",
					},
					{
						Consomation: 5,
						Begin:       "2024/05/01",
						End:         "2024/05/31",
					},
					{
						Consomation: 67,
						Begin:       "2024/10/01",
						End:         "2024/10/31",
					},
				},
			},
			expectedEstimation: []estimation.ConsomationEstimation{
				{Id: "BASE", Estimation: uint16(84)},
				{Id: "OFF-PEAK", Estimation: uint16(42)},
				{Id: "PEAK", Estimation: uint16(42)},
				{Id: "CUSTOM", Estimation: uint16(56)},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert := assert.New(t)
			estimationService := estimation.EstimationService{}
			estimation := estimationService.Estimate(&testCase.historic, tariff.TariffsRules)

			assert.Equal(testCase.expectedEstimation, estimation)
		})

	}
}
