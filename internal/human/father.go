package human

type Father struct {
}

func NewFather() Father {
	c := Father{}
	return c
}

func (c *Father) GetName() string {
	return "Tony"
}

func (c *Father) Say() string {
	return "I am " + c.GetName()
}
