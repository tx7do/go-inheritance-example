package vehicle

import "fmt"

type Train struct {
	vehicleImpl
	name string
}

func NewTrain(name string) *Train {
	c := &Train{}
	c.name = name
	c.vehicleImpl.wheelCount = 0
	return c
}

func (c *Train) GetName() string {
	return c.name
}

func (c *Train) SetName(name string) {
	c.name = name
}

func (c *Train) GetType() Type {
	return TrainType
}

func (c *Train) ToString() string {
	str := fmt.Sprintf("Train -> %s - %d", c.GetName(), c.GetWheelCount())
	return c.vehicleImpl.ToString() + str
}
