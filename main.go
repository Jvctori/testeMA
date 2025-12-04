package main

// para criar pacotes você precisa criar um diretório e passar o link ou nome dado em go mod init
import (
	"example.com/packages/ops"
	"fmt"
	// abaixo importamos uma biblioteca de terceiros (externa)
	// o comando para baixar a biblioteca é:
	// go get <link para o repósitorio do pacote>
	// isso baixará o pacote globalmente em seu sistema
	// isso adicionará ao arquivo go.mod um require com o repósitorio
	// depois precisamos importar diretamente no código
	"github.com/Pallinder/go-randomdata"
)

// func main
func main() {
	for {
		// Todas as funções usadas de pacotes devem começar com UPPERCASE
		// fmt.Scan() note que todas as func usadas da lib padrão fmt começam com UPPERCASE
		// perceba que a func PhoneNumber() começa com UPPERCASE
		fmt.Println("Random phone number:", randomdata.PhoneNumber())
		fmt.Println("Random Address:", randomdata.Address())
		ops.PresentOpt()
		ops.UserChoices()

	}

}
