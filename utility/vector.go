package utility

import (
	"errors"
	"math"
)

var (
	Z = Float64(0.0)
	O = Float64(1.0)
)

// Vector Interface
type Vector interface {
	Dim() int

	Scale(scalar Scalar) Vector

	Magnitude() float64
	Unit() Vector

	Add(vec *Vector) Vector
	Sub(vec *Vector) Vector

	Dot(vec *Vector) (float64, error)
	Cross(vec *Vector) (Vector, error)
}

// Vector Binding
type Vec2 struct{ X, Y Scalar }
type Vec3 struct{ X, Y, Z Scalar }
type Vec4 struct{ X, Y, Z, W Scalar }

func (v *Vec2) Dim() int { return 2 }
func (v *Vec3) Dim() int { return 3 }
func (v *Vec4) Dim() int { return 4 }

func (v *Vec2) Magnitude() float64 {
	return math.Sqrt(v.X.Mul(v.X).Add(v.Y.Mul(v.X)).ToFloat64())
}

func (v *Vec3) Magnitude() float64 {
	return math.Sqrt(v.X.Mul(v.X).Add(v.Y.Mul(v.X)).Add(v.Z.Mul(v.Z)).ToFloat64())
}

func (v *Vec4) Magnitude() float64 {
	return math.Sqrt(v.X.Mul(v.X).Add(v.Y.Mul(v.X)).Add(v.Z.Mul(v.Z)).Add(v.W.Mul(v.W)).ToFloat64())
}

func (v *Vec2) Unit() Vector {
	M := Float64(v.Magnitude())

	newX, errX := v.X.Div(&M)
	newY, errY := v.Y.Div(&M)

	if errX != nil || errY != nil {
		return &Vec2{X: &O, Y: &Z}
	}

	return &Vec2{X: newX, Y: newY}
}

func (v *Vec3) Unit() Vector {
	M := Float64(v.Magnitude())

	newX, errX := v.X.Div(&M)
	newY, errY := v.Y.Div(&M)
	newZ, errZ := v.Z.Div(&M)

	if errX != nil || errY != nil || errZ != nil {
		return &Vec3{X: &O, Y: &Z, Z: &Z}
	}

	return &Vec3{X: newX, Y: newY, Z: newZ}
}

func (v *Vec4) Unit() Vector {
	M := Float64(v.Magnitude())

	newX, errX := v.X.Div(&M)
	newY, errY := v.Y.Div(&M)
	newZ, errZ := v.Z.Div(&M)
	newW, errW := v.W.Div(&M)

	if errX != nil || errY != nil || errZ != nil || errW != nil {
		return &Vec4{X: &O, Y: &Z, Z: &Z, W: &Z}
	}

	return &Vec4{X: newX, Y: newY, Z: newZ, W: newW}
}

func (v *Vec2) Add(vec *Vector) Vector {
	switch V := (*vec).(type) {
	case *Vec2:
		return &Vec2{X: v.X.Add(V.X), Y: v.Y.Add(V.Y)}
	case *Vec3:
		return &Vec3{X: v.X.Add(V.X), Y: v.Y.Add(V.Y), Z: V.Z}
	case *Vec4:
		return &Vec4{X: v.X.Add(V.X), Y: v.Y.Add(V.Y), Z: V.Z, W: V.W}
	}

	return nil
}

func (v *Vec3) Add(vec *Vector) Vector {
	switch V := (*vec).(type) {
	case *Vec2:
		return &Vec3{X: v.X.Add(V.X), Y: v.Y.Add(V.Y), Z: v.Z}
	case *Vec3:
		return &Vec3{X: v.X.Add(V.X), Y: v.Y.Add(V.Y), Z: v.Z.Add(V.Z)}
	case *Vec4:
		return &Vec4{X: v.X.Add(V.X), Y: v.Y.Add(V.Y), Z: v.Z.Add(V.Z), W: V.W}
	}

	return nil
}

func (v *Vec4) Add(vec *Vector) Vector {
	switch V := (*vec).(type) {
	case *Vec2:
		return &Vec4{X: v.X.Add(V.X), Y: v.Y.Add(V.Y), Z: v.Z, W: v.W}
	case *Vec3:
		return &Vec4{X: v.X.Add(V.X), Y: v.Y.Add(V.Y), Z: v.Z.Add(V.Z), W: v.W}
	case *Vec4:
		return &Vec4{X: v.X.Add(V.X), Y: v.Y.Add(V.Y), Z: v.Z.Add(V.Z), W: v.W.Add(V.W)}
	}

	return nil
}

func (v *Vec2) Sub(vec *Vector) Vector {
	switch V := (*vec).(type) {
	case *Vec2:
		return &Vec2{X: v.X.Sub(V.X), Y: v.Y.Sub(V.Y)}
	case *Vec3:
		return &Vec3{X: v.X.Sub(V.X), Y: v.Y.Sub(V.Y), Z: V.Z}
	case *Vec4:
		return &Vec4{X: v.X.Sub(V.X), Y: v.Y.Sub(V.Y), Z: V.Z, W: V.W}
	}

	return nil
}

func (v *Vec3) Sub(vec *Vector) Vector {
	switch V := (*vec).(type) {
	case *Vec2:
		return &Vec3{X: v.X.Sub(V.X), Y: v.Y.Sub(V.Y), Z: v.Z}
	case *Vec3:
		return &Vec3{X: v.X.Sub(V.X), Y: v.Y.Sub(V.Y), Z: v.Z.Sub(V.Z)}
	case *Vec4:
		return &Vec4{X: v.X.Sub(V.X), Y: v.Y.Sub(V.Y), Z: v.Z.Sub(V.Z), W: V.W}
	}

	return nil
}

func (v *Vec4) Sub(vec *Vector) Vector {
	switch V := (*vec).(type) {
	case *Vec2:
		return &Vec4{X: v.X.Sub(V.X), Y: v.Y.Sub(V.Y), Z: v.Z, W: v.W}
	case *Vec3:
		return &Vec4{X: v.X.Sub(V.X), Y: v.Y.Sub(V.Y), Z: v.Z.Sub(V.Z), W: v.W}
	case *Vec4:
		return &Vec4{X: v.X.Sub(V.X), Y: v.Y.Sub(V.Y), Z: v.Z.Sub(V.Z), W: v.W.Sub(V.W)}
	}

	return nil
}

func (v *Vec2) Scale(scalar Scalar) Vector {
	return &Vec2{X: v.X.Mul(scalar), Y: v.Y.Mul(scalar)}
}

func (v *Vec3) Scale(scalar Scalar) Vector {
	return &Vec3{X: v.X.Mul(scalar), Y: v.Y.Mul(scalar), Z: v.Z.Mul(scalar)}
}

func (v *Vec4) Scale(scalar Scalar) Vector {
	return &Vec4{X: v.X.Mul(scalar), Y: v.Y.Mul(scalar), Z: v.Z.Mul(scalar), W: v.W.Mul(scalar)}
}

func (v *Vec2) Dot(vec *Vector) (float64, error) {
	switch t := (*vec).(type) {
	case *Vec2:
		return v.X.Mul(t.X).Add(v.Y.Mul(t.Y)).ToFloat64(), nil
	default:
		return 0.0, errors.New("cannot dot vectors with different dimensions")
	}
}

func (v *Vec3) Dot(vec *Vector) (float64, error) {
	switch t := (*vec).(type) {
	case *Vec3:
		return v.X.Mul(t.X).Add(v.Y.Mul(t.Y)).Add(v.Z.Mul(t.Z)).ToFloat64(), nil
	default:
		return 0.0, errors.New("cannot dot vectors with different dimensions")
	}
}

func (v *Vec4) Dot(vec *Vector) (float64, error) {
	switch t := (*vec).(type) {
	case *Vec4:
		return v.X.Mul(t.X).Add(v.Y.Mul(t.Y)).Add(v.Z.Mul(t.Z)).Add(v.W.Mul(t.W)).ToFloat64(), nil
	default:
		return 0.0, errors.New("cannot dot vectors with different dimensions")
	}
}

func (v *Vec2) Cross(vec *Vector) (Vector, error) {
	switch t := (*vec).(type) {
	case *Vec2:
		return &Vec3{X: &Z, Y: &Z, Z: v.X.Mul(t.Y).Sub(v.Y.Mul(t.X))}, nil
	default:
		return nil, errors.New("cannot cross vectors with different dimensions")
	}
}

func (v *Vec3) Cross(vec *Vector) (Vector, error) {
	switch t := (*vec).(type) {
	case *Vec3:
		return &Vec3{X: v.Y.Mul(t.Z).Sub(v.Z.Mul(t.Y)), Y: v.Z.Mul(t.X).Sub(v.X.Mul(t.Z)), Z: v.X.Mul(t.Y).Sub(v.Y.Mul(t.X))}, nil
	default:
		return nil, errors.New("cannot cross vectors with different dimensions")
	}
}

func (v *Vec4) Cross(vec *Vector) (Vector, error) {
	switch (*vec).(type) {
	case *Vec4:
		return nil, errors.New("cannot cross Vec4 instances")
	default:
		return nil, errors.New("cannot cross vectors with different dimensions")
	}
}
