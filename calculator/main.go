package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
	mathFunction "week-4-homework-1-damla-unal/math"
)

/*

Add error handling to the example calculator.go of ch07 we studied at the end of last class.
And also make sure that the example works file even though math function name entered by the user is lower case of name registered by the calculator.
i.e. make the calculator match math function names case-insensitive.

*/

var flag = true

// Calculator defined as a new data type
type Calculator struct {
	functions []mathFunction.MathFunction
}

// addMathFunction adds functions to function list of calculator
func (c *Calculator) addMathFunction(m mathFunction.MathFunction) {
	c.functions = append(c.functions, m)
}

// doCalculation calculate the function with arg according to name that comes as a parameter
func (c *Calculator) doCalculation(name string, arg float64) (float64, error) {
	var result float64
	for _, f := range c.functions {
		if strings.ToLower(name) == strings.ToLower(f.GetName()) {
			result = f.Calculate(arg)
			return result, nil
		}
	}
	return 0, errors.New("not found function:" + name)
}

func main() {
	//functions()
	//err := calculator()
	//if err != nil {
	//	return
	//}
	startCalculator()
}

// functions calculates some basic trigonometric functions and log function and print them
func functions() {
	sin := mathFunction.Sin{Name: "Sinus"}
	fmt.Printf("%v\n", sin)
	sin30 := sin.Calculate(math.Pi / 6)
	fmt.Printf("Sinus of 30 degree is %f\n", sin30)

	cos := mathFunction.Cos{Name: "Cosines"}
	fmt.Printf("%v\n", cos)
	cos30 := cos.Calculate(math.Pi / 6)
	fmt.Printf("Cosinus of 30 degree is %f\n", cos30)

	log := mathFunction.Log{Name: "Log of base e"}
	fmt.Printf("%v\n", log)
	logE := log.Calculate(2.71828)
	fmt.Printf("Log of Euler constant is %f\n", logE)

	var mf1 mathFunction.MathFunction = sin
	fmt.Printf("%v\n", mf1)

	mf1 = cos
	fmt.Printf("%v\n", mf1)

	mf1 = log
	fmt.Printf("%v\n", mf1)
}

// calculator calculates some basic trigonometric functions and log function and handled errors
func calculator() error {
	myCalculator := Calculator{}

	myCalculator.addMathFunction(mathFunction.Sin{Name: "Sinus"})
	myCalculator.addMathFunction(mathFunction.Cos{Name: "Cosines"})
	myCalculator.addMathFunction(mathFunction.Log{Name: "Log"})

	sin, err := myCalculator.doCalculation("Sinus", math.Pi/6)
	if err != nil {
		return err
	}
	fmt.Printf("Sinus of 30 degree is %f\n", sin)

	cos, err := myCalculator.doCalculation("Cosines", math.Pi/6)
	if err != nil {
		return err
	}
	fmt.Printf("Cosinus of 30 degree is %f\n", cos)

	log, err := myCalculator.doCalculation("Cosines", math.E)
	if err != nil {
		return err
	}
	fmt.Printf("Log of Euler constant is %f\n", log)

	return nil
}

//startCalculator takes the calculation function name and value from the user,
// then writes the calculated value to the screen.
// Error handling strategies were used while performing these operations.
func startCalculator() {
	myCalculator := Calculator{}

	myCalculator.addMathFunction(mathFunction.Sin{Name: "Sin"})
	myCalculator.addMathFunction(mathFunction.Cos{Name: "Cosine"})
	myCalculator.addMathFunction(mathFunction.Log{Name: "Log"})

	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	for flag {
		var fName string
		var arg float64
		fmt.Println("> Enter name of the calculation or enter x to exit:")
		_, err := fmt.Scanf("%s", &fName)
		if err != nil {
			fmt.Printf("Error occured while assigning the name of the calculation: %s\n", err)
			os.Exit(0)
		}

		if fName == "x" {
			//os.Exit(0)
			flag = false
		} else {
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scanf("%f", &arg)
			if err != nil {
				fmt.Printf("Error occured while assigning the value for calculation: %s\n", err)
				os.Exit(0)
			}
			value, err := myCalculator.doCalculation(fName, arg)
			if err != nil {
				fmt.Printf("Error occured: %s\n", err)
			} else {
				fmt.Printf("Result of %s of %f : %f\n", fName, arg, value)
			}
		}
	}
	println("Bye!")
}
