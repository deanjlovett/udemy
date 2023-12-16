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

//package timeparse
package main

import (
	// "errors"
	"fmt"
	"strconv"
	"strings"
)

//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components

type MyTime struct {
	hours uint
	minutes uint
	seconds uint
}

func (mt *MyTime) toString() string {
	return fmt.Sprintf("%2d:%02d:%02d",mt.hours, mt.minutes, mt.seconds)
}

func timeParse(in string) (MyTime, error){
	myTimeErr := MyTime{0,0,0}
	myTimeRet := MyTime{0,0,0}
	count := strings.Count(in,":")
	if count > 2 {
		// return myTimeErr, errors.New(fmt.Sprintf("invalid time formate: too many colons"))
		return myTimeErr, fmt.Errorf("invalid time formate: too many colons")
	}
	myArr := strings.Split(in,":")
	h, err := strconv.Atoi(myArr[0])
	if err != nil {
		return myTimeErr, err
	}else if h<0 || h>23 {
		// return myTimeErr,errors.New(fmt.Sprintf("invalid time format: the hours must bewteen 0 and 23."))
		return myTimeErr,fmt.Errorf("invalid time format: the hours must bewteen 0 and 23.")
	}
	myTimeRet.hours = uint(h)
	if len(myArr) > 1 {
		m, err := strconv.Atoi(myArr[1])
		if err != nil {
			// return myTimeErr, err
			// return myTimeErr, errors.New(fmt.Sprintf("invalid time format: can't parse minutes characters"))
			return myTimeErr, fmt.Errorf("invalid time format: can't parse minutes characters")
		}else if m<0 || m>59  || (m<10 && len(myArr[1])==1){
			// return myTimeErr,errors.New(fmt.Sprintf("invalid time format: the minutes must bewteen 0 and 59."))
			return myTimeErr, fmt.Errorf("invalid time format: the minutes must bewteen 0 and 59.")
		}	
		myTimeRet.minutes = uint(m)
	}
	if len(myArr) == 3 {
		sec, err := strconv.Atoi(myArr[2])
		if err != nil {
			// return myTimeErr, err
			// return myTimeErr, errors.New(fmt.Sprintf("invalid time format: can't parse seconds characters"))
			return myTimeErr, fmt.Errorf("invalid time format: can't parse seconds characters")
		}else if sec<0 || sec>59 || (sec<10 && len(myArr[2])==1){
			// return myTimeErr, errors.New(fmt.Sprintf("invalid time format: the seconds must bewteen 0 and 59."))
			return myTimeErr, fmt.Errorf("invalid time format: the seconds must bewteen 0 and 59.")
		}	
		myTimeRet.seconds = uint(sec)
	}
	return myTimeRet,nil
}

func main() {
	testArr :=[]string{
		"0:0:0",
		"0:00:01",
		"1:20:30",
		"25:20:30",
		"1:60:30",
		"1:20:60",
		"-1:20:59",
	}
	for i,testTime := range testArr {
		myTime, err := timeParse(testTime)
		fmt.Printf("%v: %10s, %10s, error:%v\n",i,testTime,myTime.toString(),err)

	}
}

