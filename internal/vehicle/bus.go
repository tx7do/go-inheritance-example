package vehicle

import "fmt"

type Bus struct {
	vehicleImpl
	name string
}

func NewBus(name string) *Bus {
	c := &Bus{}
	c.name = name
	c.vehicleImpl.wheelCount = 0
	return c
}

func (c *Bus) GetName() string {
	return c.name
}

func (c *Bus) SetName(name string) {
	c.name = name
}

func (c *Bus) GetType() Type {
	return BusType
}

func (c *Bus) ToString() string {
	str := fmt.Sprintf("Bus -> %s", c.GetName())
	return c.vehicleImpl.ToString() + str
}
