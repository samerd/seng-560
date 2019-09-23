package numbers

import (
	"reflect"
	"strconv"
)

//createIntNumber function creates a real number object from a int64 value
func createIntNumber(val int64) RealNumber {
	result := new(realNumber)
	result.numType = IntType
	result.intVal = val
	return result
}

//createUintNumber function creates a real number object from a uint64 value
func createUintNumber(val uint64) RealNumber {
	result := new(realNumber)
	result.numType = UintType
	result.uintVal = val
	return result
}

//createFloatNumber function creates a real number object from a float64 value
func createFloatNumber(val float64) RealNumber {
	result := new(realNumber)
	result.floatVal = val
	result.numType = FloatType
	return result
}

//createFloatNumber function creates a real number object from a float64 value
func createNumberFromString(val string) (RealNumber, error) {
	//firstly try to parse the number as an integer
	iVal, err := strconv.ParseInt(val, 10, 64)
	if err == nil {
		return createIntNumber(iVal), err
	}

	//if not, then try to parse it as float
	fVal, err := strconv.ParseFloat(val, 64)
	if err == nil {
		return createFloatNumber(fVal), err
	}

	//Otherwise, this not a numeric string
	return nil, new(NonNumericStringError)
}

//CreateNumber function creates a number object from a given value
//if the value cannot represent a number, then an error will be returned
func CreateNumber(val interface{}) (Number, error) {
	var err error
	var result RealNumber
	err = nil
	rval := reflect.ValueOf(val)
	switch val.(type) {
	case int, int8, int16, int32, int64:
		result = createIntNumber(rval.Int())
		break
	case uint, uint8, uint16, uint32, uint64:
		result = createUintNumber(rval.Uint())
		break
	case float32, float64:
		result = createFloatNumber(rval.Float())
		break
	case string:
		result, err = createNumberFromString(rval.String())
		break
	default:
		err = new(NonNumericTypeError)
		result = nil
		break
	}
	return result, err
}

//Decode a number using the given format
//Created number is always Uint
func Decode(strNum string, encoding NumberEncoding) (Number, error) {
	var base int
	switch encoding {
	case NumberEncodingHex:
		base = 16
		break
	case NumberEncodingDecimal:
		base = 10
		break
	case NumberEncodingOct:
		base = 8
		break
	case NumberEncodingBin:
		base = 2
		break
	}
	uiNum, err := strconv.ParseUint(strNum, base, 64)
	if err != nil {
		return nil, new(NonNumericStringError)
	}
	return createUintNumber(uiNum), nil
}
