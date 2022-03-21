package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntrodução() //exibeIntrodução é o bloco que não retorna nda e se encontro foro do bloco principal

	for {
		exibeMenu()
		comandoLido := leComando()

		switch comandoLido {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço seu programa")
			os.Exit(-1)
		}
	}

}

func exibeIntrodução() {
	var versão float32 = 1.1 //posso escrever assim//
	var nome = "Diego"       //posso escrever sem definir se é uma string,float ou outro//
	//idade := 35  //Ou posso simplesmente tirar tudo e colocar :=, dessa forma o golang vai interpretar//
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versão)

	//fmt.Println("O tipo da variável é", reflect.TypeOf(nome))  Para saber o tipo da variável//

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
