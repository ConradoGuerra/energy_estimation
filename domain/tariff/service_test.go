package tariff_test

import (
	"energy_estimation/domain/tariff"
	"energy_estimation/infrastructure/repositories/in_memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetTariffs(t *testing.T) {
	t.Run("it should return the list of tariffs", func(t *testing.T) {
		assert := assert.New(t)

		new_repo := in_memory.InMemoryTariffRepo{}

		tariffService := tariff.NewService(new_repo)

		gotTariff := tariffService.GetTariffs()

		assert.Equal(&[]tariff.TariffRule{
			{Id: "BASE", Ratio: 1},
			{Id: "OFF-PEAK", Ratio: .5},
			{Id: "PEAK", Ratio: .5},
			{Id: "CUSTOM", Ratio: .67},
		}, gotTariff)
	})
}
