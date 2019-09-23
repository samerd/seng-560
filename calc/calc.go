package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/samerd/seng-560/numbers"
)

//usage function describes the usage of the calc application
func usage(appPath string) {
	//extract the app name
	_, appName := filepath.Split(appPath)
	description := `Usage:
%s operand1 operator operand2
or
%s unary_operator operand
where operator: + - * / ^
and 
unary_operator: v -
`
	fmt.Printf(description, appName, appName)
}

//displayResult prints the result in decimal format and hex format if applicable
func displayResult(num numbers.Number) {
	fmt.Printf("Result = %v\n", num.GetValue())
	hex, err := num.Encode(numbers.NumberEncodingHex)
	if err == nil {
		fmt.Printf("Hex = %s\n", hex)
	}
}

type binaryFunc func(numbers.Number) (numbers.Number, error)
type unaryFunc func() (numbers.Number, error)

//binaryOperator function calculates the result of the binary operators:
//+ - * / ^
func binaryOperator(operand1, operator, operand2 string) error {
	num1, err := numbers.CreateNumber(operand1)
	if err != nil {
		return fmt.Errorf("Invalid operand #1: %s", err.Error())
	}
	num2, err := numbers.CreateNumber(operand2)
	if err != nil {
		return fmt.Errorf("Invalid operand #1: %s", err.Error())
	}

	var bFunc binaryFunc
	var nFunc string
	switch operator {
	case "+":
		bFunc = num1.Add
		nFunc = "Add"
		break
	case "-":
		bFunc = num1.Substract
		nFunc = "Substract"
	case "*":
		bFunc = num1.Multiply
		nFunc = "Multiply"
	case "/":
		bFunc = num1.Divide
		nFunc = "Divide"
		break
	case "^":
		bFunc = num1.Power
		nFunc = "Power"
		break
	default:
		return fmt.Errorf("Unsupported operator : %s", operator)
	}

	result, err := bFunc(num2)
	if err != nil {
		err = fmt.Errorf("%s error: %s", nFunc, err.Error())
	} else {
		displayResult(result)
	}
	return err
}

func unaryOperator(operator, operand string) error {
	num, err := numbers.CreateNumber(operand)
	if err != nil {
		return fmt.Errorf("Invalid operand: %s", err.Error())
	}

	var uFunc unaryFunc
	var nFunc string

	switch operator {
	case "-":
		uFunc = num.Opposite
		nFunc = "Opposite"
		break
	case "v", "V":
		uFunc = num.Sqrt
		nFunc = "Sqrt"
		break
	default:
		return fmt.Errorf("Unsupported operator : %s", operator)
	}
	result, err := uFunc()
	if err != nil {
		err = fmt.Errorf("%s error: %s", nFunc, err.Error())
	} else {
		displayResult(result)
	}
	return err
}

func main() {
	args := os.Args
	exitCode := 0
	var err error
	switch len(args) {
	case 4:
		err = binaryOperator(args[1], args[2], args[3])
		break
	case 3:
		err = unaryOperator(args[1], args[2])
		break
	default:
		usage(args[0])
		exitCode = 1
		break
	}
	if err != nil {
		fmt.Printf("-E- %s\n", err.Error())
		exitCode = 1
	}
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}
