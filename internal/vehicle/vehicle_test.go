package vehicle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBus(t *testing.T) {
	{
		var busVehicle Vehicle
		busVehicle = NewBus("big")
		assert.NotNil(t, busVehicle)

		assert.Equal(t, busVehicle.GetType(), BusType)
		assert.Equal(t, busVehicle.GetName(), "big")

		busVehicle.SetName("middle")
		assert.Equal(t, busVehicle.GetName(), "middle")
		assert.Equal(t, busVehicle.ToString(), "vehicle -> Bus -> middle")

		busVehicle.SetWheelCount(300)
		assert.Equal(t, busVehicle.GetWheelCount(), 300)
	}

	{
		bus := NewBus("little")
		assert.NotNil(t, bus)

		assert.Equal(t, bus.GetType(), BusType)
		assert.Equal(t, bus.GetName(), "little")

		assert.Equal(t, bus.ToString(), "vehicle -> Bus -> little")
	}
}

func TestTrain(t *testing.T) {
	{
		var trainVehicle Vehicle
		trainVehicle = NewTrain("big")
		assert.NotNil(t, trainVehicle)

		assert.Equal(t, trainVehicle.GetType(), TrainType)
		assert.Equal(t, trainVehicle.GetName(), "big")

		trainVehicle.SetName("middle")
		assert.Equal(t, trainVehicle.GetName(), "middle")

		assert.Equal(t, trainVehicle.GetWheelCount(), 0)
		trainVehicle.SetWheelCount(100)
		assert.Equal(t, trainVehicle.GetWheelCount(), 100)

		assert.Equal(t, trainVehicle.ToString(), "vehicle -> Train -> middle - 100")

		train := trainVehicle.(*Train)
		assert.Equal(t, train.GetWheelCount(), 100)
		assert.Equal(t, train.GetName(), "middle")
		assert.Equal(t, train.GetType(), TrainType)

		train.SetWheelCount(300)
		assert.Equal(t, train.GetWheelCount(), 300)
	}

	{
		train := NewTrain("little")
		assert.NotNil(t, train)

		assert.Equal(t, train.GetType(), TrainType)
		assert.Equal(t, train.GetName(), "little")

		assert.Equal(t, train.GetWheelCount(), 0)
		train.SetWheelCount(200)
		assert.Equal(t, train.GetWheelCount(), 200)

		assert.Equal(t, train.ToString(), "vehicle -> Train -> little - 200")

		trainVehicle := Vehicle(train)
		assert.Equal(t, trainVehicle.GetWheelCount(), 200)
		assert.Equal(t, trainVehicle.GetName(), "little")
		assert.Equal(t, trainVehicle.GetType(), TrainType)

		trainVehicle.SetWheelCount(300)
		assert.Equal(t, trainVehicle.GetWheelCount(), 300)
	}
}
