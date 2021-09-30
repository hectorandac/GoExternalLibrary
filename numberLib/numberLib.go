package numberLib

import (
	"fmt"
	"math"
)

type Number struct {
	Value float64
}

var Pi float64 = 3.1415926535

func (firstNumber Number) Add(secondNumber Number) Number {
	return Number{firstNumber.Value + secondNumber.Value}
}

func (firstNumber Number) Remove(secondNumber Number) Number {
	return Number{firstNumber.Value - secondNumber.Value}
}

func (firstNumber Number) WeirdResult(secondNumber Number) Number {
	return Number{math.Pow((firstNumber.Value*secondNumber.Value)/Pi, firstNumber.Value)}
}

func (firstNumber Number) Odd() bool {
	return math.Mod(firstNumber.Value, 2.0) != 0
}

func (firstNumber Number) Even() bool {
	return !firstNumber.Odd()
}

func (n Number) String() string {
	return fmt.Sprintf("Your value is %f", n.Value)
}
