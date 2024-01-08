package main
import "fmt"

func main(){
	var investmentAmmount = 5000.0;
	var returnRate = 5.5;
	var years = 10;

	var futureValue = investmentAmmount * returnRate * float64(years);

	fmt.Print("Future ammount is $", futureValue);
}