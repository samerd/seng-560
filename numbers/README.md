# Numbers API
## Interfaces
The numbers package defines the following interfaces:
### Number Interface
Number interface which has the following functionalities
1. Binary Operations
    1. Add - returns the addition result or an error
    1. Substract - returns the substraction result or an error
    1. Multiply - returns the multiplication result or an error
    1. Divide - returns the division result or an error
    1. Power - returns the power result or an error
1. Unary Operations
    1. Sqrt - returns the square root of the number or an error
    1. Opposite - returns the opposite of the number or an error
1. Disaply
    1. Encode - Encode the number into hex/oct/bin or return an error
    1. GetValue - returns the real value of the number

### Real Number Interface
RealNumber interface is a Number with additional functionalities:
1. Real Number attributes
    1. GetInt
    1. GetUint
    1. GetFloat
    1. AsFloat
    1. GetType

## Types
The follwoing types have been defined

### RealNumberType
This is an enumeration of the supported real type numbers:
1. IntType
1. UintType
1. FloatType

### NumberEncoding
This is an enumeration of the supported encoding types:
1. NumberEncodingHex
1. NumberEncodingDecimal
1. NumberEncodingOct
1. NumberEncodingBin

## Errors
the following errors have been implemented:

### DivisionByZeroError
DivisionByZeroError error returned in two cases:
1. in the Divide function, if the denominator is zero.
1. in the power function, if the base is zeo and the exponent is negative

### IntSignMismatchError
IntSignMismatchError error is returned when there is an operation that mixes signed and unsigned numbers

### NonRealNumberError
NonRealNumberError error is returned if the argument number is non real, in most functions, except Add and Multiply.

### MathOperationError
MathOperationError is returned when the result cannot be caculated in the real numbers range, such as square root of a negative number

### UnsupportedEncodingError
UnsupportedEncodingError is returned when trying to encode float to formats other than Decimal.

### NonNumericStringError
NonNumericStringError is returned when trying to decode a non numeric string according to the required base.

### NonNumericTypeError
NonNumericTypeError is returned when trying to make a number from an interface that does not represent a number.

## Factory
These functions supports creating Numbers

### CreateNumber
Creates a real number using general interface input, and return a RealNumber or error
The created number can be of the follwoing types
* Int
An integer number can be generated if the input is:
    * one of: int, int8, int16, int32, int64
    * string representing a decimal integer

* Uint
An unsigned integer can be generated if the input is:
    * one of uint, uint8, uint16, uint32, uint64

* Float
A Float number can be generated if the input is:
    * one of float32, float64
    * a string representing a decimal float number

All other inputs will return an error.

### Decode
Decode creates a *UintType* Number or returns an error.

The input is a string and the desired format Decimal,Hex,Oct,Bin.

# Examples
## Create a number
```golang
num, err := numbers.CreateNumber("12.5")
if err == nil {
    //Handle Error
}
```

## Add numbers
```golang
num1, err := numbers.CreateNumber("12.5")
if err == nil {
    //Handle Error
}
num2, err := numbers.CreateNumber(3)

if err == nil {
    //Handle Error
}
result, err := num1.Add(num2)
if err == nil {
    //Handle Error
}
```

## Hex Encoding
```golang
num1, err := numbers.CreateNumber(14)

if err == nil {
    //Handle Error
}
hexFmt, err := num1.Encode(numbers.NumberEncodingHex)
if err == nil {
    //Handle Error
}
fmt.Println(hexFmt)
//should print "E"
```

## Hex Decoding
```golang
hexNum := "1F"
num1, err := numbers.Decode(hexNum, numbers.NumberEncodingHex)
if err == nil {
    //Handle Error
}
decFmt, err := num1.Encode(numbers.NumberEncodingDecimal)
if err == nil {
    //Handle Error
}
fmt.Println(decFmt)
//should print "31"
```
