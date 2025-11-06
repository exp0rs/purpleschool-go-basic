package main

import (
	"errors"
	"fmt"
	"math"
)

const imtPower = 2

func main () {
	fmt.Println("Welcome to IMT calc")
	for {
		fmt.Println("Insert your height and weight")
		userHeight, userWeight := getUserInput()
		imt, err := calculateImt(userHeight,userWeight)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(imt)
		
		if !isRepeatCalculation() {
			break
		}
	}
}

func getUserInput () (userHeight float64, userWeight float64 ) {
	fmt.Print("Height in CM:")
	fmt.Scan(&userHeight)
	fmt.Print("Weght in KG:")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

func calculateImt (height float64, weight float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, errors.New("WEIGHT_OR_HEIGHT_NOT_SPECIFIED")
	}
	imt := weight / math.Pow(height/100, imtPower)
	return imt, nil
}

func outputResult (imt float64) {
	resImt := fmt.Sprintf("Your IMT: %.0f", imt)
	fmt.Println(resImt)

	switch {
	case imt < 16:
		fmt.Println("You are severely underweight.")
	case imt < 18.5:
		fmt.Println("You have a underweight.")
	case imt < 25:
		fmt.Println("You have a normal weight.")
	case imt < 30:
		fmt.Println("You have a overweight.")
	default:
		fmt.Println("You have a degree of obesity.")
	}
}

func isRepeatCalculation () bool {
	var answerYN string
	fmt.Println("Repeat calculate? y/n")
	fmt.Scan(&answerYN)
	if answerYN == "y" || answerYN == "Y" {
		return true 
	}
	return false
}