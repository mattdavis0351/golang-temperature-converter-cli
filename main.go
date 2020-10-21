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
		// Inside the for loop, below the first fmt.Print() statement which prints "What is the current temperature in...",
		// invoke the fmt.Scanln() function, passing &originValue as the argument. Assign the two return values to
		// the variables _ and err respectively. On the following line, create an if statement checking if err != nil, and if that's true,
		// invoke the printError() function, passing errReadingInput as its argument.
		fmt.Print("What is the current temperature in " + originUnit + " ? ")

		_, err := fmt.Scanln(&originValue)

		if err != nil {
			printError(errReadingInput)
		}

		if originUnit == "C" {
			convertToFahrenheit(originValue)
		} else {
			convertToCelsius(originValue)
		}

		fmt.Print("Would you like to convert another temperature ? (y/n) ")

		_, err = fmt.Scanln(&shouldConvertAgain)

		if err != nil {
			printError(errReadingInput)
		}

		if strings.ToUpper(strings.TrimSpace(shouldConvertAgain)) != "Y" {
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
