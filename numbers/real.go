package numbers

import (
	"fmt"
	"math"
)

// the realNumber struct impelments the RealNumber interfaces
// and hence the Number interfaces
type realNumber struct {
	intVal   int64
	uintVal  uint64
	floatVal float64
	numType  RealNumberType
}

//GetInt returns the interger value of the number
//valid only in case of the type is IntType
func (r *realNumber) GetInt() int64 {
	return r.intVal
}

//GetFloat returns the float value of the number
//valid only in case of the type is FloatType
func (r *realNumber) GetFloat() float64 {
	return r.floatVal
}

//GetUint returns the uint value of the number
//valid only in case of the type is UintType
func (r *realNumber) GetUint() uint64 {
	return r.uintVal
}

//GetType returns the number type
//IntType, UintType or FloatType
func (r *realNumber) GetType() RealNumberType {
	return r.numType
}

//AsFloat returns the value of the number as float, whtever its type is.
func (r *realNumber) AsFloat() float64 {
	switch r.GetType() {
	case IntType:
		return float64(r.GetInt())
	case UintType:
		return float64(r.GetUint())
	}
	return r.floatVal
}

//Add function adds two number or returns an error
//possible errors: IntSignMismatchError
func (r *realNumber) Add(num Number) (Number, error) {
	//Make sure second number is a real number
	r2, ok := num.(RealNumber)
	if !ok {
		if num == nil {
			return nil, new(NonRealNumberError)
		}
		//Real number cannot add non-real number
		// try the opposite direction
		return num.Add(r)
	}

	r1Type := r.GetType()
	r2Type := r2.GetType()

	//usigned it can be added to unsigned int
	if r1Type == UintType || r2Type == UintType {
		if r1Type != r2Type {
			return nil, new(IntSignMismatchError)
		}
		return createUintNumber(r.GetUint() + r2.GetUint()), nil
	}
	// one of the number is float the result will be float as well
	if r1Type == FloatType || r2Type == FloatType {
		return createFloatNumber(r.AsFloat() + r2.AsFloat()), nil
	}
	//the last option two integers will result in an integer
	return createIntNumber(r.GetInt() + r2.GetInt()), nil
}

//Substract function substract a number and returns the result or an error
//possible errors:
func (r *realNumber) Substract(num Number) (Number, error) {
	r2, ok := num.(RealNumber)
	if !ok {
		//cannot substract a non real number from a real number
		return nil, new(NonRealNumberError)
	}

	r1Type := r.GetType()
	r2Type := r2.GetType()

	//uint can be sustracted from uint
	if r1Type == UintType || r2Type == UintType {
		if r1Type != r2Type {
			return nil, new(IntSignMismatchError)
		}
		return createUintNumber(r.GetUint() - r2.GetUint()), nil
	}
	//if any number is a float, the result will be float
	if r1Type == FloatType || r2Type == FloatType {
		return createFloatNumber(r.AsFloat() - r2.AsFloat()), nil
	}
	//the last case is two integers, resulting in an integer
	return createIntNumber(r.GetInt() - r2.GetInt()), nil
}

//Opposite function return the opposite of the given number
// or an error in case of unsigned integer
func (r *realNumber) Opposite() (Number, error) {
	switch r.numType {
	case IntType:
		return createIntNumber(0 - r.intVal), nil
	case FloatType:
		return createFloatNumber(0 - r.floatVal), nil
	}
	return nil, new(IntSignMismatchError)
}

//Multiply function returns the multiplication result of two numbers
func (r *realNumber) Multiply(num Number) (Number, error) {
	r2, ok := num.(RealNumber)
	if !ok {
		if num == nil {
			return nil, new(NonRealNumberError)
		}
		//Real number cannot add non-real number
		// try the opposite direction
		return num.Multiply(r)
	}

	r1Type := r.GetType()
	r2Type := r2.GetType()

	//take care of signed/unsigned mismatch
	if r1Type == UintType || r2Type == UintType {
		if r1Type != r2Type {
			return nil, new(IntSignMismatchError)
		}
		return createUintNumber(r.GetUint() * r2.GetUint()), nil
	}

	//If any number i of float type, return a float number
	if r1Type == FloatType || r2Type == FloatType {
		return createFloatNumber(r.AsFloat() * r2.AsFloat()), nil
	}
	//if the two numbers are integers, return an integer
	return createIntNumber(r.GetInt() * r2.GetInt()), nil
}

//Divide function returns the divison reulr of the two numbers or an error.
// possible errors are: NonRealNumberError, DivisionByZeroError or IntSignMismatchError
func (r *realNumber) Divide(num Number) (Number, error) {
	r2, ok := num.(RealNumber)
	if !ok {
		// cannot divide a non real number
		return nil, new(NonRealNumberError)
	}

	//prevent dividing by zero
	if r2.AsFloat() == 0 {
		return nil, new(DivisionByZeroError)
	}

	r1Type := r.GetType()
	r2Type := r2.GetType()

	//take care of signed/unsigned mismatch
	if r1Type == UintType || r2Type == UintType {
		if r1Type != r2Type {
			return nil, new(IntSignMismatchError)
		}
		return createUintNumber(r.GetUint() / r2.GetUint()), nil
	}

	//If any number is float return a float, otherwise return an integer
	if r1Type == FloatType || r2Type == FloatType {
		return createFloatNumber(r.AsFloat() / r2.AsFloat()), nil
	}
	return createIntNumber(r.GetInt() / r2.GetInt()), nil
}

//Power function returns the power result or an error
//Result if exists, will be float
func (r *realNumber) Power(num Number) (Number, error) {
	r2, ok := num.(RealNumber)
	if !ok {
		return nil, new(NonRealNumberError)
	}

	base := r.AsFloat()
	exponent := r2.AsFloat()
	if base == 0 && exponent < 0 {
		return nil, new(DivisionByZeroError)
	}

	result := math.Pow(base, exponent)
	if math.IsNaN(result) {
		return nil, new(MathOperationError)
	}

	return createFloatNumber(result), nil
}

//Sqrt function return the square roor of a number or an error
//result is always a float or a math error
func (r *realNumber) Sqrt() (Number, error) {
	val := r.AsFloat()
	result := math.Sqrt(val)
	if math.IsNaN(result) {
		return nil, new(MathOperationError)
	}
	return createFloatNumber(result), nil
}

//GetValue function returns the value of the number according to its type
func (r *realNumber) GetValue() interface{} {
	var val interface{}
	switch r.GetType() {
	case FloatType:
		val = r.GetFloat()
		break
	case IntType:
		val = r.GetInt()
		break
	case UintType:
		val = r.GetUint()
		break
	}
	return val
}

//Encode function returns the encoded string using the given format
func (r *realNumber) Encode(encoding NumberEncoding) (string, error) {
	var res string

	//float does not support non-decimal encoding
	if r.GetType() == FloatType && encoding != NumberEncodingDecimal {
		return "", new(UnsupportedEncodingError)
	}

	val := r.GetValue()

	switch encoding {
	case NumberEncodingDecimal:
		res = fmt.Sprintf("%v", val)
		break
	case NumberEncodingHex:
		res = fmt.Sprintf("%x", val)
		break
	case NumberEncodingOct:
		res = fmt.Sprintf("%o", val)
		break
	case NumberEncodingBin:
		res = fmt.Sprintf("%b", val)
		break
	}
	return res, nil
}
