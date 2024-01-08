package main

import (
	"fmt"
	"math"
)

func main(){
	const inflationRate float64 = 4;
	var investmentAmmount, years float64;
	
	//shorthand to declare and assign new variable
	returnRate := 5.5;

	fmt.Print("Investment ammount: ");
	fmt.Scan(&investmentAmmount);

	fmt.Print("Years: ");
	fmt.Scan(&years);

	futureValue := investmentAmmount * math.Pow((1 + returnRate / 100), years);
	//round 3 decimal places
	futureValue = math.Round(futureValue * 1000) / 1000;
	fmt.Println("Future ammount is $", futureValue);

	futureRealValue := futureValue / math.Pow((1 + inflationRate/100), years);
	futureRealValue = math.Round(futureRealValue*1000)/1000;
	fmt.Println("Future ammount with 4% inflation is $", futureRealValue);
}