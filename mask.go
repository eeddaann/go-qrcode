package qrcode

import (
	"github.com/yeqown/go-qrcode/v2/matrix"
)

// maskPatternModulo ...
// mask Pattern ref to: https://www.thonky.com/qr-code-tutorial/mask-patterns
type maskPatternModulo uint32

const (
	// modulo0 (x+y) mod 2 == 0
	modulo0 maskPatternModulo = iota
	// modulo1 (x) mod 2 == 0
	modulo1
	// modulo2 (y) mod 3 == 0
	modulo2
	// modulo3 (x+y) mod 3 == 0
	modulo3
	// modulo4 (floor (x/ 2) + floor (y/ 3) mod 2 == 0
	modulo4
	// modulo5 (x * y) mod 2) + (x * y) mod 3) == 0
	modulo5
	// modulo6 (x * y) mod 2) + (x * y) mod 3) mod 2 == 0
	modulo6
	// modulo7 (x + y) mod 2) + (x * y) mod 3) mod 2 == 0
	modulo7
)

type mask struct {
	mat      *matrix.Matrix    // matrix
	mode     maskPatternModulo // mode
	moduloFn moduloFunc        // moduloFn masking function
}

// newMask ...
func newMask(mat *matrix.Matrix, mode maskPatternModulo) *mask {
	m := &mask{
		mat:      mat.Copy(),
		mode:     mode,
		moduloFn: getModuloFunc(mode),
	}
	m.masking()

	return m
}

// moduloFunc to define what's modulo func
type moduloFunc func(int, int) bool

func getModuloFunc(mode maskPatternModulo) (f moduloFunc) {
	f = nil
	switch mode {
	case modulo0:
		f = modulo0Func
	case modulo1:
		f = modulo1Func
	case modulo2:
		f = modulo2Func
	case modulo3:
		f = modulo3Func
	case modulo4:
		f = modulo4Func
	case modulo5:
		f = modulo5Func
	case modulo6:
		f = modulo6Func
	case modulo7:
		f = modulo7Func
	}

	return
}

// init generate maks by mode
func (m *mask) masking() {
	moduloFn := m.moduloFn
	if moduloFn == nil {
		panic("impossible panic, contact maintainer plz")
	}

	m.mat.Iterate(matrix.COLUMN, func(x, y int, s matrix.State) {
		// skip the function modules
		if state, _ := m.mat.Get(x, y); state != matrix.StateInit {
			_ = m.mat.Set(x, y, matrix.StateInit)
			return
		}
		if moduloFn(x, y) {
			_ = m.mat.Set(x, y, matrix.StateTrue)
		} else {
			_ = m.mat.Set(x, y, matrix.StateFalse)
		}
	})
}

// modulo0Func for maskPattern function
// modulo0 (x+y) mod 2 == 0
func modulo0Func(x, y int) bool {
	return (x+y)%2 == 0
}

// modulo1Func for maskPattern function
// modulo1 (y) mod 2 == 0
func modulo1Func(x, y int) bool {
	return y%2 == 0
}

// modulo2Func for maskPattern function
// modulo2 (x) mod 3 == 0
func modulo2Func(x, y int) bool {
	return x%3 == 0
}

// modulo3Func for maskPattern function
// modulo3 (x+y) mod 3 == 0
func modulo3Func(x, y int) bool {
	return (x+y)%3 == 0
}

// modulo4Func for maskPattern function
// modulo4 (floor (x/ 2) + floor (y/ 3) mod 2 == 0
func modulo4Func(x, y int) bool {
	return (x/3+y/2)%2 == 0
}

// modulo5Func for maskPattern function
// modulo5 (x * y) mod 2 + (x * y) mod 3 == 0
func modulo5Func(x, y int) bool {
	return (x*y)%2+(x*y)%3 == 0
}

// modulo6Func for maskPattern function
// modulo6 (x * y) mod 2) + (x * y) mod 3) mod 2 == 0
func modulo6Func(x, y int) bool {
	return ((x*y)%2+(x*y)%3)%2 == 0
}

// modulo7Func for maskPattern function
// modulo7 (x + y) mod 2) + (x * y) mod 3) mod 2 == 0
func modulo7Func(x, y int) bool {
	return ((x+y)%2+(x*y)%3)%2 == 0
}
