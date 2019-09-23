package numbers

//NumberEncoding - enumeration type of number format
type NumberEncoding int

//Supported enums of number format
const (
	NumberEncodingDecimal NumberEncoding = iota
	NumberEncodingHex
	NumberEncodingOct
	NumberEncodingBin
)

//Number interface
type Number interface {
	Add(num Number) (Number, error)
	Substract(num Number) (Number, error)
	Opposite() (Number, error)
	Multiply(num Number) (Number, error)
	Divide(num Number) (Number, error)
	Sqrt() (Number, error)
	Power(Number) (Number, error)
	Encode(NumberEncoding) (string, error)
	GetValue() interface{}
}

//RealNumber interface represents a real number
//real number is a number that can be flot, int or uint
type RealNumber interface {
	Number
	GetFloat() float64
	GetInt() int64
	GetUint() uint64
	AsFloat() float64
	GetType() RealNumberType
}

//RealNumberType enumeration of real number types
type RealNumberType int

//real number types
const (
	IntType RealNumberType = iota
	UintType
	FloatType
)
