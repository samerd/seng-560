package numbers

//DivisionByZeroError struct represent an error return in two cases
//1. Divide by zero.
//2. Power of zero to a negative exponent
type DivisionByZeroError struct {
}

//Error function return the description of the DivisionByZeroError error
func (e *DivisionByZeroError) Error() string {
	return "Division by zero"
}

//IntSignMismatchError struct represent an error return in case
//we have an operation between usigned interger and another signed type.
type IntSignMismatchError struct {
}

//Error function return the description of the IntSignMismatchError error
func (e *IntSignMismatchError) Error() string {
	return "Signed/Unsigned mismatch"
}

//NonRealNumberError struct represent an error return in case
//we have an operation between real number and non real number
type NonRealNumberError struct {
}

//Error function return the description of the NonRealNumberError error
func (e *NonRealNumberError) Error() string {
	return "Not a real number"
}

//MathOperationError struct represents an error returned in case
//we have an error while calculaing powers
type MathOperationError struct {
}

//Error function return the description of the MathOperationError error
func (e *MathOperationError) Error() string {
	return "Not a real number"
}

//UnsupportedEncodingError struct represents an error returned in case
//there is an encoding request for non-decimal float
type UnsupportedEncodingError struct {
}

//Error function return the description of the UnsupportedEncodingError error
func (e *UnsupportedEncodingError) Error() string {
	return "Encoding is not supported for this number type"
}

//NonNumericStringError struct represents an error returned in case
//of decoding a non numeric string
type NonNumericStringError struct {
}

//Error function return the description of the NonNumericStringError error
func (e *NonNumericStringError) Error() string {
	return "Invalid Numeric literals"
}

//NonNumericTypeError struct represents an error returned in case
//of decoding a non numeric string
type NonNumericTypeError struct {
}

//Error function return the description of the NonNumericTypeError error
func (e *NonNumericTypeError) Error() string {
	return "Non Numeric Type"
}
