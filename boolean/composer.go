package boolean

type Composer struct {
	composers []BoolComposer
	invoker   BoolComposer
}

func NewComposer(composers ...BoolComposer) *Composer {
	return &Composer{
		composers: composers,
		invoker:   Compose(composers...),
	}
}

func (c *Composer) Invoke(first, second bool) bool {
	return c.invoker(first, second)
}

func (c *Composer) Append(composer BoolComposer) *Composer {
	c.composers = append(c.composers, composer)
	c.invoker = Compose(c.composers...)
	return c
}

func (c *Composer) Prepend(composer BoolComposer) *Composer {
	composers := make([]BoolComposer, 0, len(c.composers)+1)
	composers[0] = composer
	c.composers = append(composers, c.composers...)
	c.invoker = Compose(c.composers...)
	return c
}

func (c *Composer) Insert(at int, composer BoolComposer) *Composer {
	composers := make([]BoolComposer, len(c.composers)+1)
	composers[at] = composer

	for i, comp := range c.composers[0:at] {
		composers[i] = comp
	}

	for i, comp := range c.composers[at+1:] {
		composers[i] = comp
	}

	c.composers = composers
	c.invoker = Compose(c.composers...)
	return c
}
