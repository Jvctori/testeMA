package ops

import (
	"fmt"
	"os"
)

func UserChoices() {

	var choice int

	accountBalance, err := GetFloatFromFile()
	if err != nil {
		fmt.Println("Não foi possível acessar sua conta!")
	}

	fmt.Println("Digite uma opção:")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		fmt.Println("Your balance is:", accountBalance)

	case 2:
		var depositAmount float64
		fmt.Println("The value you want to deposit is:")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		WriteFloatInFile(accountBalance)

	case 3:
		var whitdrawlAmount float64
		fmt.Println("The amount you want to withdral:")
		fmt.Scan(&whitdrawlAmount)
		accountBalance -= whitdrawlAmount
		WriteFloatInFile(accountBalance)

	default:
		fmt.Println("Thx for using our bank")
		WriteFloatInFile(accountBalance)
		os.Exit(0)

	}
}
