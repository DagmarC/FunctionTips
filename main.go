package main

import (
	"FunctionsDeepDive/controlFlow"
	"fmt"
	"math"
)

type MathExpression1 = string

const (
	AddExpression      = MathExpression1("add")
	SubtractExpression = MathExpression1("subtract")
	MultiplyExpression = MathExpression1("multiply")
)

func main() {
	addExpression := mathExpression()
	fmt.Println(addExpression(1, 3))

	// Call and return multiple functions
	subtractExpr := mathExpression1(SubtractExpression)
	fmt.Println(subtractExpr(1, 3))

	fmt.Println(mathExpression1(MultiplyExpression)(2, 3))

	// Call function with function
	// 2*(3+2)
	fmt.Println("Function called with function: ", double(3, 2, mathExpression1(AddExpression)))

	// Holds state and remembers value.
	powerOfTwo := powerOfTwo()
	fmt.Println(powerOfTwo())
	// Call again and x will be incremented from 2 to 3 and returned its power of two.
	fmt.Println(powerOfTwo())

	fmt.Println("----CONTROL FLOW----")
	// Functions - Control Flow
	//err1 := controlFlow.ErrorHandling()
	//if err1 != nil {
	//	return
	//}

	err := controlFlow.ReadFullFile()
	if err != nil {
		println("Main: ", err.Error())
	}

}

// returns function
func mathExpression() func(float64, float64) float64 {
	return func(f float64, f2 float64) float64 {
		return f + f2
	}
}

// returns function
func mathExpression1(expression MathExpression1) func(float64, float64) float64 {
	return func(f float64, f2 float64) float64 {
		switch expression {
		case AddExpression:
			return f + f2
		case MultiplyExpression:
			return f * f2
		case SubtractExpression:
			return f - f2
		default:
			return 0
		}
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}

// Hold state within one function for another function to use.
// Each time I call this function frm the main I will increment x by 1.
// x starts at 1 and increments to two and the it returns 4.
// Next call: x increments to three and returns 9, ...
func powerOfTwo() func() int64 {
	x := 1.0
	// Modify var from the outside state.
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}
