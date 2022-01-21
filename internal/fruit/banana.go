package fruit

type Banana struct {
	name   string
	energy float32
}

func NewBanana(name string) *Banana {
	c := &Banana{}
	c.name = name
	c.energy = 0
	return c
}

func (c *Banana) GetName() string {
	return c.name
}

func (c *Banana) SetName(name string) {
	c.name = name
}

func (c *Banana) GetType() Type {
	return BananaType
}

func (c *Banana) SetEnergy(energy float32) {
	c.energy = energy
}

func (c *Banana) GetEnergy() float32 {
	return c.energy
}
