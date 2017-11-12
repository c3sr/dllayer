package layer

type Base struct {
	name string
}

func (c Base) Name() string {
	return c.name
}

func (c *Base) SetName(s string) {
	c.name = s
}
