package human

type Mother struct {
}

func NewMother() Mother {
	c := Mother{}
	return c
}

func (c *Mother) GetName() string {
	return "Aurora"
}

func (c *Mother) Say() string {
	return "I am " + c.GetName()
}
