package vehicle

type Type int

const (
	BusType   Type = 0
	TrainType Type = 1
)

type Vehicle interface {
	GetType() Type

	GetName() string
	SetName(name string)

	SetWheelCount(wheelCount int)
	GetWheelCount() int

	ToString() string
}

type vehicleImpl struct {
	wheelCount int
}

func (c *vehicleImpl) SetWheelCount(wheelCount int) {
	c.wheelCount = wheelCount
}

func (c *vehicleImpl) GetWheelCount() int {
	return c.wheelCount
}

func (c *vehicleImpl) ToString() string {
	return "vehicle -> "
}
