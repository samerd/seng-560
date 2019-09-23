package numbers

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	inputs := [][3]interface{}{
		{1, 2, (int64)(3)},
		{1, -1, (int64)(0)},
		{(uint64)(5), (uint64)(5), (uint64)(10)}}
	for index, input := range inputs {
		num1, _ := CreateNumber(input[0])
		num2, _ := CreateNumber(input[1])
		num3, err := num1.Add(num2)
		if err != nil {
			t.Errorf("Add failed on input #%d", index)
		}
		if num3.GetValue() != input[2] {
			t.Errorf("Result error on add input #%d", index)
		}
	}
	val1 := (float64)(1.1)
	val2 := -1
	res := (float64)(0.1)
	num1, _ := CreateNumber(val1)
	num2, _ := CreateNumber(val2)
	num3, err := num1.Add(num2)
	if err != nil {
		t.Errorf("Add failed on float input")
	}
	diff := math.Abs(num3.GetValue().(float64) - res)
	if diff >= 1e-9 {
		t.Errorf("Result error on add float")
	}
}

func TestAddError(t *testing.T) {
	num, _ := CreateNumber(1)
	var num2 Number
	_, err := num.Add(num2)
	if err == nil {
		t.Error("Adding nil to a number should fail")
	}

	num2, _ = CreateNumber(uint(2))
	_, err = num.Add(num2)
	if err == nil {
		t.Error("Adding int and uint should fail")
	}
}

func TestSubstract(t *testing.T) {
	inputs := [][3]interface{}{
		{1, 2, (int64)(-1)},
		{1, -1, (int64)(2)},
		{(uint64)(5), (uint64)(5), (uint64)(0)}}
	for index, input := range inputs {
		num1, _ := CreateNumber(input[0])
		num2, _ := CreateNumber(input[1])
		num3, err := num1.Substract(num2)
		if err != nil {
			t.Errorf("Substract failed on input #%d", index)
		}
		if num3.GetValue() != input[2] {
			t.Errorf("Result error on Substract input #%d", index)
		}
	}
	val1 := (float64)(1.1)
	val2 := -1
	res := (float64)(2.1)
	num1, _ := CreateNumber(val1)
	num2, _ := CreateNumber(val2)
	num3, err := num1.Substract(num2)
	if err != nil {
		t.Errorf("Substract failed on float input")
	}
	diff := math.Abs(num3.GetValue().(float64) - res)
	if diff >= 1e-9 {
		t.Errorf("Result error on substract float")
	}
}

func TestSubstractError(t *testing.T) {
	num, _ := CreateNumber(1)
	var num2 Number
	_, err := num.Substract(num2)
	if err == nil {
		t.Error("Substract nil to a number should fail")
	}

	num2, _ = CreateNumber(uint(2))
	_, err = num.Substract(num2)
	if err == nil {
		t.Error("Substract int and uint should fail")
	}
}

func TestMultiply(t *testing.T) {
	inputs := [][3]interface{}{
		{1, 2, (int64)(2)},
		{1, -1, (int64)(-1)},
		{(uint64)(5), (uint64)(5), (uint64)(25)}}
	for index, input := range inputs {
		num1, _ := CreateNumber(input[0])
		num2, _ := CreateNumber(input[1])
		num3, err := num1.Multiply(num2)
		if err != nil {
			t.Errorf("Multiply failed on input #%d", index)
		}
		if num3.GetValue() != input[2] {
			t.Errorf("Result error on Multiply input #%d", index)
		}
	}
	val1 := (float64)(1.1)
	val2 := -1
	res := (float64)(-1.1)
	num1, _ := CreateNumber(val1)
	num2, _ := CreateNumber(val2)
	num3, err := num1.Multiply(num2)
	if err != nil {
		t.Errorf("Multiply failed on float input")
	}
	diff := math.Abs(num3.GetValue().(float64) - res)
	if diff >= 1e-9 {
		t.Errorf("Result error on Multiply float")
	}
}

func TestMultiplyError(t *testing.T) {
	num, _ := CreateNumber(1)
	var num2 Number
	_, err := num.Multiply(num2)
	if err == nil {
		t.Error("Multiply nil to a number should fail")
	}

	num2, _ = CreateNumber(uint(2))
	_, err = num.Multiply(num2)
	if err == nil {
		t.Error("Multiply int and uint should fail")
	}
}

func TestDivide(t *testing.T) {
	inputs := [][3]interface{}{
		{1, 2, (int64)(0)},
		{1, -1, (int64)(-1)},
		{(uint64)(5), (uint64)(5), (uint64)(1)}}
	for index, input := range inputs {
		num1, _ := CreateNumber(input[0])
		num2, _ := CreateNumber(input[1])
		num3, err := num1.Divide(num2)
		if err != nil {
			t.Errorf("Divide failed on input #%d", index)
		}
		if num3.GetValue() != input[2] {
			t.Errorf("Result error on Divide input #%d", index)
		}
	}
	val1 := (float64)(1.1)
	val2 := -1
	res := (float64)(-1.1)
	num1, _ := CreateNumber(val1)
	num2, _ := CreateNumber(val2)
	num3, err := num1.Divide(num2)
	if err != nil {
		t.Errorf("Divide failed on float input")
	}
	diff := math.Abs(num3.GetValue().(float64) - res)
	if diff >= 1e-9 {
		t.Errorf("Result error on Divide float")
	}
}

func TestDivideError(t *testing.T) {
	num, _ := CreateNumber(1)
	var num2 Number
	_, err := num.Divide(num2)
	if err == nil {
		t.Error("Divide nil to a number should fail")
	}

	num2, _ = CreateNumber(uint(2))
	_, err = num.Divide(num2)
	if err == nil {
		t.Error("Divide int and uint should fail")
	}

	num2, _ = CreateNumber(0)
	_, err = num.Divide(num2)
	if err == nil {
		t.Error("Divide by zero should fail")
	}
}

func TestPower(t *testing.T) {
	inputs := [][3]interface{}{
		{2, 2, (float64)(4)},
		{4, -1, (float64)(.25)},
		{(uint64)(5), (uint64)(2), (float64)(25)},
		{4, .5, (float64)(2.0)}}
	for index, input := range inputs {
		num1, _ := CreateNumber(input[0])
		num2, _ := CreateNumber(input[1])
		num3, err := num1.Power(num2)
		if err != nil {
			t.Errorf("Power failed on input #%d", index)
		}
		res := input[2].(float64)
		diff := math.Abs(num3.GetValue().(float64) - res)
		if diff >= 1e-9 {
			t.Errorf("Result error on Power input #%d", index)
		}
	}
}

func TestPowerError(t *testing.T) {
	num, _ := CreateNumber(-1)
	var num2 Number
	_, err := num.Power(num2)
	if err == nil {
		t.Error("Power nil for a number should fail")
	}

	num2, _ = CreateNumber(0)
	_, err = num2.Power(num)
	if err == nil {
		t.Error("Power by zero should fail")
	}
}

func TestSqrt(t *testing.T) {
	inputs := [][2]interface{}{
		{4, (float64)(2)},
		{9, (float64)(3)},
		{(uint64)(5), math.Sqrt(5)}}

	for index, input := range inputs {
		num1, _ := CreateNumber(input[0])
		res := input[1].(float64)
		num2, err := num1.Sqrt()
		if err != nil {
			t.Errorf("Sqrt failed on input #%d", index)
		}
		diff := math.Abs(num2.GetValue().(float64) - res)
		if diff >= 1e-9 {
			t.Errorf("Result error on Sqrt input #%d", index)
		}
	}
}

func TestSqrtError(t *testing.T) {
	num, _ := CreateNumber(-4)
	_, err := num.Sqrt()
	if err == nil {
		t.Error("Sqrt should fail on negative input")
	}
}

func TestOpposite(t *testing.T) {
	inputs := [][2]interface{}{
		{4, (int64)(-4)},
		{-9.1, (float64)(9.1)},
		{0, (int64)(0)}}

	for index, input := range inputs {
		num1, _ := CreateNumber(input[0])
		num2, err := num1.Opposite()
		if err != nil {
			t.Errorf("Opposite failed on input #%d", index)
		}
		if num2.GetValue() != input[1] {
			t.Errorf("Result error on Opposite input #%d", index)
		}
	}
}

func TestOppositeError(t *testing.T) {
	num, _ := CreateNumber((uint64)(4))
	_, err := num.Opposite()
	if err == nil {
		t.Error("Sqrt should fail on sign mismatch input")
	}
}

func TestEncode(t *testing.T) {
	inputs := [][3]interface{}{
		{4, NumberEncodingBin, "100"},
		{16, NumberEncodingHex, "10"},
		{12, NumberEncodingOct, "14"},
		{(float64)(20), NumberEncodingDecimal, "20"}}

	for index, input := range inputs {
		num1, _ := CreateNumber(input[0])
		strNum, err := num1.Encode(input[1].(NumberEncoding))
		if err != nil {
			t.Errorf("Encode failed on input #%d", index)
		}
		if strNum != input[2] {
			t.Errorf("Result error on Encode input #%d", index)
		}
	}
}

func TestEncodeError(t *testing.T) {
	num, _ := CreateNumber((float64)(4))
	_, err := num.Encode(NumberEncodingHex)
	if err == nil {
		t.Error("Encode should fail on float input")
	}
}

func TestGetValue(t *testing.T) {
	inputs := [][2]interface{}{
		{4, (int64)(4)},
		{(uint)(5), (uint64)(5)},
		{-1.1, (float64)(-1.1)}}

	for index, input := range inputs {
		num1, _ := CreateNumber(input[0])
		val := num1.GetValue()
		if val != input[1] {
			t.Errorf("Result error on GetValue input #%d", index)
		}
	}
}
