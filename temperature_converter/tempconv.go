package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

/*
Exercise:
	Create an abstraction to make conversions among different temperature systems such as Celsius, Fahrenheit and Kelvin.
	Create a calculator that is loosely coupled to that abstraction
	i.e. program to an interface, not an implementation principles is observed.
	Make sure that the program can interact with the user so that user can achieve some conversions using the program.
*/
var flag = true

//Used a combination of an interface (abstract interface) and struct (abstract concrete type).
//Together they can provide the functionalities of an abstract class.

// Define an interface as achieve abstraction
type Calculator interface {
	convert(to string) (float64, error)
}

// Define a new data type "Celsius"
//Celsius implements the calculator interface methods
type Celsius struct {
	degree float64
}

func (c Celsius) GetCDegree() float64 {
	return c.degree
}

// implement convert method
func (c Celsius) convert(to string) (float64, error) {
	var degree float64
	var err error
	switch to {
	case "C":
		degree = c.degree
	case "K":
		degree = c.degree + 273.15
	case "F":
		degree = (9 * c.degree / 5) + 32
	default:
		err = errors.New("invalid temperature name")
	}

	return degree, err
}

// Define a new data type "Fahrenheit"
//Fahrenheit implements the calculator interface methods
type Fahrenheit struct {
	degree float64
}

func (f Fahrenheit) GetFDegree() float64 {
	return f.degree
}

// implement convert method
func (f Fahrenheit) convert(to string) (float64, error) {
	var degree float64
	var err error
	switch to {
	case "C":
		degree = (f.degree - 32) * 5 / 9
	case "K":
		degree = (f.degree-32)*5/9 + 273.15
	case "F":
		degree = f.degree
	default:
		err = errors.New("invalid temperature name")
	}

	return degree, err
}

// Define a new data type "Kelvin"
//Kelvin implements the calculator interface methods
type Kelvin struct {
	degree float64
}

func (k Kelvin) GetKDegree() float64 {
	return k.degree
}

// implement convert method
func (k Kelvin) convert(to string) (float64, error) {
	var degree float64
	var err error
	switch to {
	case "C":
		degree = k.degree - 273.15
	case "K":
		degree = k.degree
	case "F":
		degree = (k.degree-273.15)*9/5 + 32
	default:
		err = errors.New("invalid temperature name")
	}

	return degree, err
}

func main() {

	for flag {
		var tName, to string
		var degreeV float64

		fmt.Println("> Enter the value of the temperature: ")
		_, err := fmt.Scanf("%f", &degreeV)
		if err != nil {
			fmt.Printf("Error occured while assigning the value of the temperature: %s\n", err)
			os.Exit(0)
		}

		fmt.Println("> Enter abbreviation of the temperature or enter x to exit: ")
		_, err = fmt.Scanf("%s", &tName)
		if err != nil {
			fmt.Printf("Error occured while assigning the name of the temperature: %s\n", err)
			os.Exit(0)
		}

		if tName == "x" {
			flag = false
		} else {
			fmt.Println("> Enter the temperature abbreviation you want to convert: ")
			_, err = fmt.Scanf("%s", &to)
			to = strings.ToUpper(to)
			if err != nil {
				fmt.Printf("Error occured while assigning the name of the requested: %s\n", err)
				os.Exit(0)
			}

			switch strings.ToUpper(tName) {
			case "C":
				c := Celsius{
					degree: degreeV,
				}
				converted, err := c.convert(to)
				if err != nil {
					return
				}
				fmt.Printf("From %f Celsius to %s => %f\n", degreeV, to, converted)
			case "K":
				k := Kelvin{
					degree: degreeV,
				}
				converted, err := k.convert(to)
				if err != nil {
					return
				}
				fmt.Printf("From %f Kelvin to %s => %f\n", degreeV, to, converted)

			case "F":
				f := Fahrenheit{
					degree: degreeV,
				}
				converted, err := f.convert(to)
				if err != nil {
					return
				}
				fmt.Printf("From %f Fahrenheit to %s => %f\n", degreeV, to, converted)

			default:
				fmt.Println("Invalid temperature type")
				os.Exit(0)

			}
		}
	}

}
