package main

import "fmt"

import "math"

func main(){
	var investmentAmmount = 5000;
	var returnRate = 5.5;
	
	//shorthand to declare and assign new variable
	years := 10;

	futureValue := float64(investmentAmmount) * math.Pow((1 + returnRate / 100), float64(years));

	//round 3 decimal places
	futureValue = math.Round(futureValue * 1000) / 1000;

	fmt.Print("Future ammount is $", futureValue);
}