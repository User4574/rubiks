package cube

import (
	"errors"
	"strings"
)

type colour uint8

const (
	White colour = iota
	Yellow
	Red
	Orange
	Blue
	Green
)

func (c colour) String() string {
	switch c {
	case White:
		return "[97mW[0m"
	case Yellow:
		return "[97mY[0m"
	case Red:
		return "[97mR[0m"
	case Orange:
		return "[97mO[0m"
	case Blue:
		return "[97mB[0m"
	case Green:
		return "[97mG[0m"
	}
	return ""
}

type row []colour

func (r row) String() (s string) {
	for i := range r {
		s += r[i].String()
	}
	return
}

type face []row

func (f face) colGet(c int) (r row) {
	r = make(row, len(f))
	for i := 0; i < len(f); i++ {
		r[i] = f[i][c]
	}
	return
}

func (f *face) colSet(c int, r row) {
	for i := 0; i < len(*f); i++ {
		(*f)[i][c] = r[i]
	}
}

type Cube struct {
	front  face
	back   face
	top    face
	bottom face
	left   face
	right  face
}

func New(size int) (*Cube, error) {
	if size < 1 {
		return nil, errors.New("Size must be greater than 0.")
	}
	var c Cube
	c.front = make(face, size)
	c.back = make(face, size)
	c.top = make(face, size)
	c.bottom = make(face, size)
	c.left = make(face, size)
	c.right = make(face, size)

	for i := 0; i < size; i++ {
		c.front[i] = make(row, size)
		c.back[i] = make(row, size)
		c.top[i] = make(row, size)
		c.bottom[i] = make(row, size)
		c.left[i] = make(row, size)
		c.right[i] = make(row, size)

		for j := 0; j < size; j++ {
			c.front[i][j] = Red
			c.back[i][j] = Orange
			c.top[i][j] = White
			c.bottom[i][j] = Yellow
			c.left[i][j] = Green
			c.right[i][j] = Blue
		}
	}

	return &c, nil
}

func (c Cube) Size() int {
	return len(c.front)
}

func (c Cube) String() (r string) {
	s := strings.Repeat(" ", c.Size()+2)

	for i := 0; i < c.Size(); i++ {
		r += s + c.top[i].String() + "\n"
	}
	r += "\n"

	for i := 0; i < c.Size(); i++ {
		r += c.left[i].String() + "  " +
			c.front[i].String() + "  " +
			c.right[i].String() + "  " +
			c.back[i].String() + "\n"
	}
	r += "\n"

	for i := 0; i < c.Size(); i++ {
		r += s + c.bottom[i].String() + "\n"
	}

	return
}
