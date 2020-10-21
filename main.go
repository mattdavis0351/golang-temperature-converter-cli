package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var originUnit string
var originValue float64

var shouldConvertAgain string

var err error

var errInvalidArguments = errors.New("Invalid arguments")
var errReadingInput = errors.New("Error reading input")

func main() {
	// Find the main() function definition. On the very first line inside this function,
	// use the built-in len function to check if the length of os.Args is different than 2.
	// If so, invoke the printError() function, passing errInvalidArguments as the argument.
	if len(os.Args) != 2 {
		printError(errInvalidArguments)
	}

	// Below the if statement we just wrote, invoke the strings.ToUpper() function passing os.Args[1] as the argument.
	// This ensures consistency when reading command line arguments provided by the user.
	// Assign the result to the previously defined originUnit variable.
	originUnit = strings.ToUpper(os.Args[1])

	for {
		fmt.Print("What is the current temperature in " + originUnit + " ? ")

		fmt.Print("Would you like to convert another temperature ? (y/n) ")

		if shouldConvertAgain != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}
