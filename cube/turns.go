package cube

func (a *row) flip() {
	for left, right := 0, len(*a)-1; left < right; left, right = left+1, right-1 {
		(*a)[left], (*a)[right] = (*a)[right], (*a)[left]
	}
}

func (f *face) transpose() {
	l := len(*f)
	nf := make(face, l)

	for i := 0; i < l; i++ {
		nf[i] = make(row, l)
		for j := 0; j < l; j++ {
			nf[i][j] = (*f)[j][i]
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			(*f)[i][j] = nf[i][j]
		}
	}
}

func (f *face) clockwise() {
	f.transpose()

	for i := 0; i < len(*f); i++ {
		(*f)[i].flip()
	}
}

func (f *face) _clockwise() {
	for i := 0; i < len(*f); i++ {
		(*f)[i].flip()
	}

	f.transpose()
}

func (c *Cube) Y(layer int) {
	_layer := c.Size() - layer - 1

	if layer == 0 {
		c.top.clockwise()
	} else if _layer == 0 {
		c.bottom._clockwise()
	}

	temp := c.front[layer]
	c.front[layer] = c.right[layer]
	c.right[layer] = c.back[layer]
	c.back[layer] = c.left[layer]
	c.left[layer] = temp
}

func (c *Cube) Y_(layer int) {
	_layer := c.Size() - layer - 1

	if layer == 0 {
		c.top._clockwise()
	} else if _layer == 0 {
		c.bottom.clockwise()
	}

	temp := c.front[layer]
	c.front[layer] = c.left[layer]
	c.left[layer] = c.back[layer]
	c.back[layer] = c.right[layer]
	c.right[layer] = temp
}

func (c *Cube) U() {
	c.Y(0)
}
func (c *Cube) U_() {
	c.Y_(0)
}

func (c *Cube) D() {
	c.Y_(c.Size() - 1)
}
func (c *Cube) D_() {
	c.Y(c.Size() - 1)
}

func (c *Cube) Z(layer int) {
	_layer := c.Size() - layer - 1

	if layer == 0 {
		c.front.clockwise()
	} else if _layer == 0 {
		c.back._clockwise()
	}

	temp := c.top[_layer]
	c.top[_layer] = c.left.colGet(_layer)
	c.top[_layer].flip()
	c.left.colSet(_layer, c.bottom[layer])
	c.bottom[layer] = c.right.colGet(layer)
	c.bottom[layer].flip()
	c.right.colSet(layer, temp)
}

func (c *Cube) Z_(layer int) {
	_layer := c.Size() - layer - 1

	if layer == 0 {
		c.front._clockwise()
	} else if _layer == 0 {
		c.back.clockwise()
	}

	temp := c.top[_layer]
	c.top[_layer] = c.right.colGet(layer)
	c.bottom[layer].flip()
	c.right.colSet(layer, c.bottom[layer])
	c.bottom[layer] = c.left.colGet(layer)
	temp.flip()
	c.left.colSet(_layer, temp)
}

func (c *Cube) F() {
	c.Z(0)
}
func (c *Cube) F_() {
	c.Z_(0)
}

func (c *Cube) B() {
	c.Z_(c.Size() - 1)
}
func (c *Cube) B_() {
	c.Z(c.Size() - 1)
}

func (c *Cube) X(layer int) {
	_layer := c.Size() - layer - 1

	if layer == 0 {
		c.right.clockwise()
	} else if _layer == 0 {
		c.left._clockwise()
	}

	temp := c.top.colGet(_layer)
	temp.flip()
	c.top.colSet(_layer, c.front.colGet(_layer))
	c.front.colSet(_layer, c.bottom.colGet(_layer))
	t := c.back.colGet(layer)
	t.flip()
	c.bottom.colSet(_layer, t)
	c.back.colSet(layer, temp)
}

func (c *Cube) X_(layer int) {
	_layer := c.Size() - layer - 1

	if layer == 0 {
		c.right._clockwise()
	} else if _layer == 0 {
		c.left.clockwise()
	}

	temp := c.top.colGet(_layer)
	t := c.back.colGet(layer)
	t.flip()
	c.top.colSet(_layer, t)
	t = c.bottom.colGet(_layer)
	t.flip()
	c.back.colSet(layer, t)
	c.bottom.colSet(_layer, c.front.colGet(_layer))
	c.front.colSet(_layer, temp)
}

func (c *Cube) R() {
	c.X(0)
}
func (c *Cube) R_() {
	c.X_(0)
}

func (c *Cube) L() {
	c.X_(c.Size() - 1)
}
func (c *Cube) L_() {
	c.X(c.Size() - 1)
}
