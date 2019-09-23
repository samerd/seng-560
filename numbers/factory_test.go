package numbers

import (
	"testing"
)

func TestCreateNumber(t *testing.T) {
	inputs := [][2]interface{}{
		{1, IntType},
		{uint(1), UintType},
		{1.0, FloatType},
		{"1", IntType},
		{"1.0", FloatType}}
	for index, input := range inputs {
		num, err := CreateNumber(input[0])
		if err != nil {
			t.Errorf("Create number failed for input #%d", index)
		}
		rnum, ok := num.(RealNumber)
		if !ok {
			t.Errorf("Got a non real number for input #%d", index)
		}
		if rnum.GetType() != input[1] {
			t.Errorf("Got a wrong number type for input #%d", index)
		}
	}
}

func TestCreateNumberError(t *testing.T) {
	inputs := []interface{}{
		"1.s", ([]byte)("1.0"), []int{1, 2}}
	for index, input := range inputs {
		num, err := CreateNumber(input)
		if err == nil || num != nil {
			t.Errorf("create number shoudl fail for input #%d", index)
		}
	}
}

func TestCreateInt(t *testing.T) {
	val := int64(123)
	num := createIntNumber(val)
	if num.GetType() != IntType {
		t.Error("Got wrong number type instead of int")
	}
	if num.GetInt() != val {
		t.Error("Got wrong integer value")
	}
}

func TestCreateUint(t *testing.T) {
	val := uint64(123)
	num := createUintNumber(val)
	if num.GetType() != UintType {
		t.Error("Got wrong number type instead of unsigned int")
	}
	if num.GetUint() != val {
		t.Error("Got wrong unsigned integer value")
	}
}

func TestCreateFloat(t *testing.T) {
	val := 1.2
	num := createFloatNumber(val)
	if num.GetType() != FloatType {
		t.Error("Got wrong number type instead of float")
	}
	if num.GetFloat() != val {
		t.Error("Got wrong float value")
	}
}

func TestCreateFromString(t *testing.T) {
	val := "1"
	num, _ := createNumberFromString(val)
	if num.GetType() != IntType {
		t.Error("Got wrong number type from string instead of int")
	}
	if num.GetInt() != 1 {
		t.Error("Got wrong int value from string")
	}

	val = "1.0"
	num, _ = createNumberFromString(val)
	if num.GetType() != FloatType {
		t.Error("Got wrong number type from string instead of float")
	}
	if num.GetFloat() != 1.0 {
		t.Error("Got wrong float value from string")
	}
}

func TestCreateFromStringError(t *testing.T) {
	val := "1F"
	_, err := createNumberFromString(val)
	if err == nil {
		t.Error("createNumberFromString should fail on invalid numeric string")
	}
}

func TestDecode(t *testing.T) {
	inputs := [][3]interface{}{
		{"12", NumberEncodingDecimal, 12},
		{"12", NumberEncodingHex, 18},
		{"12", NumberEncodingOct, 10},
		{"11", NumberEncodingBin, 3}}
	for index, input := range inputs {
		strNum := input[0].(string)
		encoding := input[1].(NumberEncoding)
		val := uint64(input[2].(int))
		num, err := Decode(strNum, encoding)
		if err != nil {
			t.Errorf("Failed to decode input #%d", index)
		}
		rnum, ok := num.(RealNumber)
		if !ok {
			t.Errorf("Non real number in input #%d", index)
		}
		if rnum.GetUint() != val {
			t.Errorf("value mismatch in input #%d", index)
		}
	}
}

func TestDecodeError(t *testing.T) {
	inputs := [][2]interface{}{
		{"12", NumberEncodingBin},
		{"9", NumberEncodingOct},
		{"1F", NumberEncodingDecimal},
		{"GG", NumberEncodingHex}}
	for index, input := range inputs {
		strNum := input[0].(string)
		encoding := input[1].(NumberEncoding)
		_, err := Decode(strNum, encoding)
		if err == nil {
			t.Errorf("Decode should fail for input #%d", index)
		}
	}
}
