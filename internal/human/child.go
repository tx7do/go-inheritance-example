package human

type Child struct {
	Father
	*Mother
}

func NewChild() Child {
	c := Child{}
	return c
}

func (c *Child) GetName() string {
	return "Jerry"
}

func (c *Child) Say() string {
	return "I am " + c.GetName() + ", My Father is " + c.Father.Say() + ", My Mother is " + c.Mother.Say() + "."
}
