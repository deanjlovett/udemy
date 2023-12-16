//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

//package main

import (
	"fmt"
	"strconv"
	"strings"
)

//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components

type MyTime struct {
	hours, minutes, seconds int
}

func (mt *MyTime) toString() string {
	return fmt.Sprintf("%2d:%02d:%02d", mt.hours, mt.minutes, mt.seconds)
}

type MyTimeParseError struct {
	str   string
	input string
}

func (t *MyTimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.str, t.input)
}

func TimeParse(input string) (MyTime, error) {

	myTimeErr := MyTime{0, 0, 0}

	myArr := strings.Split(input, ":")
	count := len(myArr)
	if count < 2 || count > 3 {
		// return myTimeErr, errors.New(fmt.Sprintf("invalid time formate: too many colons"))
		// return myTimeErr, fmt.Errorf("invalid time formate: too many colons")
		return myTimeErr, &MyTimeParseError{"Invalid number of time components", input}
	}

	//* If parsing fails, then a descriptive error must be returned

	hours, err := strconv.Atoi(myArr[0])
	if err != nil {
		return myTimeErr, &MyTimeParseError{fmt.Sprintf("Error parsing hours: %v", err), input}
	}

	minutes, err := strconv.Atoi(myArr[1])
	if err != nil {
		return myTimeErr, &MyTimeParseError{fmt.Sprintf("Error parsing minutes: %v", err), input}
	}

	seconds := 0
	if count == 3 {
		seconds, err = strconv.Atoi(myArr[2])
		if err != nil {
			return myTimeErr, &MyTimeParseError{fmt.Sprintf("Error parsing seconds: %v", err), input}
		}
	}

	if hours < 0 || hours > 23 {
		return myTimeErr, &MyTimeParseError{"hours out of range: 0 <= hour <= 23", fmt.Sprintf("%v", hours)}
	}
	if minutes > 59 || minutes < 0 {
		return myTimeErr, &MyTimeParseError{"minutes out of range: 0 <= minute <= 59", fmt.Sprintf("%v", minutes)}
	}
	if seconds > 59 || seconds < 0 {
		return myTimeErr, &MyTimeParseError{"seconds out of range: 0 <= second <= 59", fmt.Sprintf("%v", seconds)}
	}
	return MyTime{hours, minutes, seconds}, nil
}
