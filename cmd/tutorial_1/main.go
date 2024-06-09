package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main(){
	var intNumb int = 18;
	fmt.Println(intNumb)

	var intFloat float32 = 34435.6
	fmt.Println(intFloat)

	var myString string = `Ashraf 
	John`

	fmt.Println(myString)
	fmt.Println(len(myString))
	fmt.Println(utf8.RuneCountInString(myString))

	var1, var2 := 1, "23"
	fmt.Println(var1, var2)

	printMe()
	var numerator int = 50
	var denominator int = 5
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf(`The answer is %v and %v`, result, remainder)
}

func printMe () {
	fmt.Println("First Function")
}

func intDivision (numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}