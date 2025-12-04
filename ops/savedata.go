package ops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// func que pega o saldo da conta de um arquivo txt
const accountBalanceFile = "balance.txt"

func GetFloatFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	// se houver algum valor em err que representa um error
	// exibirá uma mansagem no terminal
	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	// vai pegar o arquivo do txt e converter para string
	balanceText := string(data)

	// depois vai converter essa string em um float 64
	// strconv.ParseFloat(string, bitsize(tamanho do tipo do dado float64 ou float32))
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

// func que salva o saldo no arquivo .txt
func WriteFloatInFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	// os.WriteFile()
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}
