package utility

import (
	"errors"
)

// Scalar Interface
type Scalar interface {
	Integral() bool

	ToInt() int
	ToInt8() int8
	ToInt16() int16
	ToInt32() int32
	ToInt64() int64
	ToFloat32() float32
	ToFloat64() float64

	Add(Scalar) Scalar
	Sub(Scalar) Scalar
	Mul(Scalar) Scalar
	Div(Scalar) (Scalar, error)
	Mod(Scalar) (Scalar, error)
}

// Integer Scalar Binding
type Int int
type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64

// Float Scalar Binding
type Float32 float32
type Float64 float64

func (i *Int) Integral() bool     { return true }
func (i *Int8) Integral() bool    { return true }
func (i *Int16) Integral() bool   { return true }
func (i *Int32) Integral() bool   { return true }
func (i *Int64) Integral() bool   { return true }
func (f *Float32) Integral() bool { return false }
func (f *Float64) Integral() bool { return false }

func (i Int) ToInt() int     { return int(i) }
func (i Int8) ToInt() int    { return int(i) }
func (i Int16) ToInt() int   { return int(i) }
func (i Int32) ToInt() int   { return int(i) }
func (i Int64) ToInt() int   { return int(i) }
func (f Float32) ToInt() int { return int(f) }
func (f Float64) ToInt() int { return int(f) }

func (i Int) ToInt8() int8     { return int8(i) }
func (i Int8) ToInt8() int8    { return int8(i) }
func (i Int16) ToInt8() int8   { return int8(i) }
func (i Int32) ToInt8() int8   { return int8(i) }
func (i Int64) ToInt8() int8   { return int8(i) }
func (f Float32) ToInt8() int8 { return int8(f) }
func (f Float64) ToInt8() int8 { return int8(f) }

func (i Int) ToInt16() int16     { return int16(i) }
func (i Int8) ToInt16() int16    { return int16(i) }
func (i Int16) ToInt16() int16   { return int16(i) }
func (i Int32) ToInt16() int16   { return int16(i) }
func (i Int64) ToInt16() int16   { return int16(i) }
func (f Float32) ToInt16() int16 { return int16(f) }
func (f Float64) ToInt16() int16 { return int16(f) }

func (i Int) ToInt32() int32     { return int32(i) }
func (i Int8) ToInt32() int32    { return int32(i) }
func (i Int16) ToInt32() int32   { return int32(i) }
func (i Int32) ToInt32() int32   { return int32(i) }
func (i Int64) ToInt32() int32   { return int32(i) }
func (f Float32) ToInt32() int32 { return int32(f) }
func (f Float64) ToInt32() int32 { return int32(f) }

func (i Int) ToInt64() int64     { return int64(i) }
func (i Int8) ToInt64() int64    { return int64(i) }
func (i Int16) ToInt64() int64   { return int64(i) }
func (i Int32) ToInt64() int64   { return int64(i) }
func (i Int64) ToInt64() int64   { return int64(i) }
func (f Float32) ToInt64() int64 { return int64(f) }
func (f Float64) ToInt64() int64 { return int64(f) }

func (i Int) ToFloat32() float32     { return float32(i) }
func (i Int8) ToFloat32() float32    { return float32(i) }
func (i Int16) ToFloat32() float32   { return float32(i) }
func (i Int32) ToFloat32() float32   { return float32(i) }
func (i Int64) ToFloat32() float32   { return float32(i) }
func (f Float32) ToFloat32() float32 { return float32(f) }
func (f Float64) ToFloat32() float32 { return float32(f) }

func (i Int) ToFloat64() float64     { return float64(i) }
func (i Int8) ToFloat64() float64    { return float64(i) }
func (i Int16) ToFloat64() float64   { return float64(i) }
func (i Int32) ToFloat64() float64   { return float64(i) }
func (i Int64) ToFloat64() float64   { return float64(i) }
func (f Float32) ToFloat64() float64 { return float64(f) }
func (f Float64) ToFloat64() float64 { return float64(f) }

func (i *Int) Add(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) + Int(*s)
		return &result
	case *Int8:
		result := (*i) + Int(*s)
		return &result
	case *Int16:
		result := (*i) + Int(*s)
		return &result
	case *Int32:
		result := (*i) + Int(*s)
		return &result
	case *Int64:
		result := (*i) + Int(*s)
		return &result
	case *Float32:
		result := (*i) + Int(*s)
		return &result
	case *Float64:
		result := (*i) + Int(*s)
		return &result
	}

	return nil
}

func (i *Int8) Add(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) + Int8(*s)
		return &result
	case *Int8:
		result := (*i) + Int8(*s)
		return &result
	case *Int16:
		result := (*i) + Int8(*s)
		return &result
	case *Int32:
		result := (*i) + Int8(*s)
		return &result
	case *Int64:
		result := (*i) + Int8(*s)
		return &result
	case *Float32:
		result := (*i) + Int8(*s)
		return &result
	case *Float64:
		result := (*i) + Int8(*s)
		return &result
	}

	return nil
}

func (i *Int16) Add(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) + Int16(*s)
		return &result
	case *Int8:
		result := (*i) + Int16(*s)
		return &result
	case *Int16:
		result := (*i) + Int16(*s)
		return &result
	case *Int32:
		result := (*i) + Int16(*s)
		return &result
	case *Int64:
		result := (*i) + Int16(*s)
		return &result
	case *Float32:
		result := (*i) + Int16(*s)
		return &result
	case *Float64:
		result := (*i) + Int16(*s)
		return &result
	}

	return nil
}

func (i *Int32) Add(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) + Int32(*s)
		return &result
	case *Int8:
		result := (*i) + Int32(*s)
		return &result
	case *Int16:
		result := (*i) + Int32(*s)
		return &result
	case *Int32:
		result := (*i) + Int32(*s)
		return &result
	case *Int64:
		result := (*i) + Int32(*s)
		return &result
	case *Float32:
		result := (*i) + Int32(*s)
		return &result
	case *Float64:
		result := (*i) + Int32(*s)
		return &result
	}

	return nil
}

func (i *Int64) Add(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) + Int64(*s)
		return &result
	case *Int8:
		result := (*i) + Int64(*s)
		return &result
	case *Int16:
		result := (*i) + Int64(*s)
		return &result
	case *Int32:
		result := (*i) + Int64(*s)
		return &result
	case *Int64:
		result := (*i) + Int64(*s)
		return &result
	case *Float32:
		result := (*i) + Int64(*s)
		return &result
	case *Float64:
		result := (*i) + Int64(*s)
		return &result
	}

	return nil
}

func (f *Float32) Add(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*f) + Float32(*s)
		return &result
	case *Int8:
		result := (*f) + Float32(*s)
		return &result
	case *Int16:
		result := (*f) + Float32(*s)
		return &result
	case *Int32:
		result := (*f) + Float32(*s)
		return &result
	case *Int64:
		result := (*f) + Float32(*s)
		return &result
	case *Float32:
		result := (*f) + Float32(*s)
		return &result
	case *Float64:
		result := (*f) + Float32(*s)
		return &result
	}

	return nil
}

func (f *Float64) Add(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*f) + Float64(*s)
		return &result
	case *Int8:
		result := (*f) + Float64(*s)
		return &result
	case *Int16:
		result := (*f) + Float64(*s)
		return &result
	case *Int32:
		result := (*f) + Float64(*s)
		return &result
	case *Int64:
		result := (*f) + Float64(*s)
		return &result
	case *Float32:
		result := (*f) + Float64(*s)
		return &result
	case *Float64:
		result := (*f) + Float64(*s)
		return &result
	}

	return nil
}

func (i *Int) Sub(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) - Int(*s)
		return &result
	case *Int8:
		result := (*i) - Int(*s)
		return &result
	case *Int16:
		result := (*i) - Int(*s)
		return &result
	case *Int32:
		result := (*i) - Int(*s)
		return &result
	case *Int64:
		result := (*i) - Int(*s)
		return &result
	case *Float32:
		result := (*i) - Int(*s)
		return &result
	case *Float64:
		result := (*i) - Int(*s)
		return &result
	}

	return nil
}

func (i *Int8) Sub(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) - Int8(*s)
		return &result
	case *Int8:
		result := (*i) - Int8(*s)
		return &result
	case *Int16:
		result := (*i) - Int8(*s)
		return &result
	case *Int32:
		result := (*i) - Int8(*s)
		return &result
	case *Int64:
		result := (*i) - Int8(*s)
		return &result
	case *Float32:
		result := (*i) - Int8(*s)
		return &result
	case *Float64:
		result := (*i) - Int8(*s)
		return &result
	}

	return nil
}

func (i *Int16) Sub(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) - Int16(*s)
		return &result
	case *Int8:
		result := (*i) - Int16(*s)
		return &result
	case *Int16:
		result := (*i) - Int16(*s)
		return &result
	case *Int32:
		result := (*i) - Int16(*s)
		return &result
	case *Int64:
		result := (*i) - Int16(*s)
		return &result
	case *Float32:
		result := (*i) - Int16(*s)
		return &result
	case *Float64:
		result := (*i) - Int16(*s)
		return &result
	}

	return nil
}

func (i *Int32) Sub(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) - Int32(*s)
		return &result
	case *Int8:
		result := (*i) - Int32(*s)
		return &result
	case *Int16:
		result := (*i) - Int32(*s)
		return &result
	case *Int32:
		result := (*i) - Int32(*s)
		return &result
	case *Int64:
		result := (*i) - Int32(*s)
		return &result
	case *Float32:
		result := (*i) - Int32(*s)
		return &result
	case *Float64:
		result := (*i) - Int32(*s)
		return &result
	}

	return nil
}

func (i *Int64) Sub(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) - Int64(*s)
		return &result
	case *Int8:
		result := (*i) - Int64(*s)
		return &result
	case *Int16:
		result := (*i) - Int64(*s)
		return &result
	case *Int32:
		result := (*i) - Int64(*s)
		return &result
	case *Int64:
		result := (*i) - Int64(*s)
		return &result
	case *Float32:
		result := (*i) - Int64(*s)
		return &result
	case *Float64:
		result := (*i) - Int64(*s)
		return &result
	}

	return nil
}

func (f *Float32) Sub(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*f) - Float32(*s)
		return &result
	case *Int8:
		result := (*f) - Float32(*s)
		return &result
	case *Int16:
		result := (*f) - Float32(*s)
		return &result
	case *Int32:
		result := (*f) - Float32(*s)
		return &result
	case *Int64:
		result := (*f) - Float32(*s)
		return &result
	case *Float32:
		result := (*f) - Float32(*s)
		return &result
	case *Float64:
		result := (*f) - Float32(*s)
		return &result
	}

	return nil
}

func (f *Float64) Sub(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*f) - Float64(*s)
		return &result
	case *Int8:
		result := (*f) - Float64(*s)
		return &result
	case *Int16:
		result := (*f) - Float64(*s)
		return &result
	case *Int32:
		result := (*f) - Float64(*s)
		return &result
	case *Int64:
		result := (*f) - Float64(*s)
		return &result
	case *Float32:
		result := (*f) - Float64(*s)
		return &result
	case *Float64:
		result := (*f) - Float64(*s)
		return &result
	}

	return nil
}

func (i *Int) Mul(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) * Int(*s)
		return &result
	case *Int8:
		result := (*i) * Int(*s)
		return &result
	case *Int16:
		result := (*i) * Int(*s)
		return &result
	case *Int32:
		result := (*i) * Int(*s)
		return &result
	case *Int64:
		result := (*i) * Int(*s)
		return &result
	case *Float32:
		result := (*i) * Int(*s)
		return &result
	case *Float64:
		result := (*i) * Int(*s)
		return &result
	}

	return nil
}

func (i *Int8) Mul(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) * Int8(*s)
		return &result
	case *Int8:
		result := (*i) * Int8(*s)
		return &result
	case *Int16:
		result := (*i) * Int8(*s)
		return &result
	case *Int32:
		result := (*i) * Int8(*s)
		return &result
	case *Int64:
		result := (*i) * Int8(*s)
		return &result
	case *Float32:
		result := (*i) * Int8(*s)
		return &result
	case *Float64:
		result := (*i) * Int8(*s)
		return &result
	}

	return nil
}

func (i *Int16) Mul(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) * Int16(*s)
		return &result
	case *Int8:
		result := (*i) * Int16(*s)
		return &result
	case *Int16:
		result := (*i) * Int16(*s)
		return &result
	case *Int32:
		result := (*i) * Int16(*s)
		return &result
	case *Int64:
		result := (*i) * Int16(*s)
		return &result
	case *Float32:
		result := (*i) * Int16(*s)
		return &result
	case *Float64:
		result := (*i) * Int16(*s)
		return &result
	}

	return nil
}

func (i *Int32) Mul(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) * Int32(*s)
		return &result
	case *Int8:
		result := (*i) * Int32(*s)
		return &result
	case *Int16:
		result := (*i) * Int32(*s)
		return &result
	case *Int32:
		result := (*i) * Int32(*s)
		return &result
	case *Int64:
		result := (*i) * Int32(*s)
		return &result
	case *Float32:
		result := (*i) * Int32(*s)
		return &result
	case *Float64:
		result := (*i) * Int32(*s)
		return &result
	}

	return nil
}

func (i *Int64) Mul(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*i) * Int64(*s)
		return &result
	case *Int8:
		result := (*i) * Int64(*s)
		return &result
	case *Int16:
		result := (*i) * Int64(*s)
		return &result
	case *Int32:
		result := (*i) * Int64(*s)
		return &result
	case *Int64:
		result := (*i) * Int64(*s)
		return &result
	case *Float32:
		result := (*i) * Int64(*s)
		return &result
	case *Float64:
		result := (*i) * Int64(*s)
		return &result
	}

	return nil
}

func (f *Float32) Mul(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*f) * Float32(*s)
		return &result
	case *Int8:
		result := (*f) * Float32(*s)
		return &result
	case *Int16:
		result := (*f) * Float32(*s)
		return &result
	case *Int32:
		result := (*f) * Float32(*s)
		return &result
	case *Int64:
		result := (*f) * Float32(*s)
		return &result
	case *Float32:
		result := (*f) * Float32(*s)
		return &result
	case *Float64:
		result := (*f) * Float32(*s)
		return &result
	}

	return nil
}

func (f *Float64) Mul(scalar Scalar) Scalar {
	switch s := scalar.(type) {
	case *Int:
		result := (*f) * Float64(*s)
		return &result
	case *Int8:
		result := (*f) * Float64(*s)
		return &result
	case *Int16:
		result := (*f) * Float64(*s)
		return &result
	case *Int32:
		result := (*f) * Float64(*s)
		return &result
	case *Int64:
		result := (*f) * Float64(*s)
		return &result
	case *Float32:
		result := (*f) * Float64(*s)
		return &result
	case *Float64:
		result := (*f) * Float64(*s)
		return &result
	}

	return nil
}

func (i *Int) Div(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Int(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int(*s)
		return &result, nil
	case *Int8:
		if Int(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int(*s)
		return &result, nil
	case *Int16:
		if Int(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int(*s)
		return &result, nil
	case *Int32:
		if Int(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int(*s)
		return &result, nil
	case *Int64:
		if Int(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int(*s)
		return &result, nil
	case *Float32:
		if Int(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int(*s)
		return &result, nil
	case *Float64:
		if Int(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (i *Int8) Div(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Int8(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int8(*s)
		return &result, nil
	case *Int8:
		if Int8(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int8(*s)
		return &result, nil
	case *Int16:
		if Int8(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int8(*s)
		return &result, nil
	case *Int32:
		if Int8(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int8(*s)
		return &result, nil
	case *Int64:
		if Int8(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int8(*s)
		return &result, nil
	case *Float32:
		if Int8(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int8(*s)
		return &result, nil
	case *Float64:
		if Int8(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int8(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (i *Int16) Div(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Int16(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int16(*s)
		return &result, nil
	case *Int8:
		if Int16(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int16(*s)
		return &result, nil
	case *Int16:
		if Int16(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int16(*s)
		return &result, nil
	case *Int32:
		if Int16(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int16(*s)
		return &result, nil
	case *Int64:
		if Int16(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int16(*s)
		return &result, nil
	case *Float32:
		if Int16(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int16(*s)
		return &result, nil
	case *Float64:
		if Int16(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int16(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (i *Int32) Div(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Int32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int32(*s)
		return &result, nil
	case *Int8:
		if Int32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int32(*s)
		return &result, nil
	case *Int16:
		if Int32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int32(*s)
		return &result, nil
	case *Int32:
		if Int32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int32(*s)
		return &result, nil
	case *Int64:
		if Int32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int32(*s)
		return &result, nil
	case *Float32:
		if Int32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int32(*s)
		return &result, nil
	case *Float64:
		if Int32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int32(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (i *Int64) Div(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Int64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int64(*s)
		return &result, nil
	case *Int8:
		if Int64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int64(*s)
		return &result, nil
	case *Int16:
		if Int64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int64(*s)
		return &result, nil
	case *Int32:
		if Int64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int64(*s)
		return &result, nil
	case *Int64:
		if Int64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int64(*s)
		return &result, nil
	case *Float32:
		if Int64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int64(*s)
		return &result, nil
	case *Float64:
		if Int64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*i) / Int64(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (f *Float32) Div(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Int8:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Int16:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Int32:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Int64:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Float32:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Float64:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (f *Float64) Div(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Int8:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Int16:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Int32:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Int64:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Float32:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Float64:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (i *Int) Mod(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Int(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int(*s)
		return &result, nil
	case *Int8:
		if Int(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int(*s)
		return &result, nil
	case *Int16:
		if Int(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int(*s)
		return &result, nil
	case *Int32:
		if Int(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int(*s)
		return &result, nil
	case *Int64:
		if Int(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int(*s)
		return &result, nil
	case *Float32:
		if Int(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int(*s)
		return &result, nil
	case *Float64:
		if Int(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (i *Int8) Mod(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Int8(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int8(*s)
		return &result, nil
	case *Int8:
		if Int8(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int8(*s)
		return &result, nil
	case *Int16:
		if Int8(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int8(*s)
		return &result, nil
	case *Int32:
		if Int8(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int8(*s)
		return &result, nil
	case *Int64:
		if Int8(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int8(*s)
		return &result, nil
	case *Float32:
		if Int8(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int8(*s)
		return &result, nil
	case *Float64:
		if Int8(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int8(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (i *Int16) Mod(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Int16(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int16(*s)
		return &result, nil
	case *Int8:
		if Int16(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int16(*s)
		return &result, nil
	case *Int16:
		if Int16(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int16(*s)
		return &result, nil
	case *Int32:
		if Int16(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int16(*s)
		return &result, nil
	case *Int64:
		if Int16(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int16(*s)
		return &result, nil
	case *Float32:
		if Int16(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int16(*s)
		return &result, nil
	case *Float64:
		if Int16(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int16(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (i *Int32) Mod(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Int32(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int32(*s)
		return &result, nil
	case *Int8:
		if Int32(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int32(*s)
		return &result, nil
	case *Int16:
		if Int32(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int32(*s)
		return &result, nil
	case *Int32:
		if Int32(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int32(*s)
		return &result, nil
	case *Int64:
		if Int32(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int32(*s)
		return &result, nil
	case *Float32:
		if Int32(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int32(*s)
		return &result, nil
	case *Float64:
		if Int32(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int32(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (i *Int64) Mod(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Int64(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int64(*s)
		return &result, nil
	case *Int8:
		if Int64(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int64(*s)
		return &result, nil
	case *Int16:
		if Int64(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int64(*s)
		return &result, nil
	case *Int32:
		if Int64(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int64(*s)
		return &result, nil
	case *Int64:
		if Int64(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int64(*s)
		return &result, nil
	case *Float32:
		if Int64(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int64(*s)
		return &result, nil
	case *Float64:
		if Int64(*s) == 0 {
			return nil, errors.New("mod by 0")
		}
		result := (*i) % Int64(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (f *Float32) Mod(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Int8:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Int16:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Int32:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Int64:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Float32:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	case *Float64:
		if Float32(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float32(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}

func (f *Float64) Mod(scalar Scalar) (Scalar, error) {
	switch s := scalar.(type) {
	case *Int:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Int8:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Int16:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Int32:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Int64:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Float32:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	case *Float64:
		if Float64(*s) == 0 {
			return nil, errors.New("division by 0")
		}
		result := (*f) / Float64(*s)
		return &result, nil
	}

	return nil, errors.New("unhandled Scalar type")
}
