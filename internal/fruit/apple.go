package fruit

type Apple struct {
	name string
}

func NewApple(name string) *Apple {
	c := &Apple{}
	c.name = name
	return c
}

func (c *Apple) GetName() string {
	return c.name
}

func (c *Apple) SetName(name string) {
	c.name = name
}

func (c *Apple) GetType() Type {
	return AppleType
}
