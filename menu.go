package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

)


type Menu struct {
	help int
	cadastrarNota int
	consultarNota int
	consultarNotas int
	consultarCR int
	terminar int
}

func imprimirTexto(texto string) {
	fmt.Println(texto)
}

func (m *Menu) configurarTeclas() {
	m.help = 0
	m.cadastrarNota = 1
	m.consultarNota = 2
	m.consultarNotas = 3
	m.consultarCR = 4
	m.terminar = 9
}

func capturarEntrada(nomeEntrada string) string {
	imprimirTexto("\nDigite um(a) " + nomeEntrada + ":")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entrada := scanner.Text()
	return entrada
	
}
	

func (m Menu) ExibirOpcoes() {
	imprimirTexto("\n*****\nAtalhos de teclado:\n")
	imprimirTexto("Cadastrar Nota = "+ strconv.Itoa(m.cadastrarNota))
	imprimirTexto("Consultar Nota = "+ strconv.Itoa(m.consultarNota))
	imprimirTexto("Consultar Notas = "+ strconv.Itoa(m.consultarNotas))
	imprimirTexto("Consultar CR = "+ strconv.Itoa(m.consultarCR))
	imprimirTexto("Exibir os atalhos novamente = "+ strconv.Itoa(m.help))
	imprimirTexto("Terminar a execução do programa = "+ strconv.Itoa(m.terminar))
	imprimirTexto("\n*****\n")



}


func (m Menu) ExecutarAcoes() {
	
	scanner := bufio.NewScanner(os.Stdin)
	for  scanner.Scan(){

		//entrada:= scanner.Scan()
		escolha,erro := strconv.Atoi(scanner.Text())


		if erro != nil {
		  fmt.Errorf("Erro: Valor informado não é um número!")
		  fmt.Println(erro)
		}


		switch escolha {
			case m.help:
			  	m.ExibirOpcoes()
			  	break

			case m.cadastrarNota:
			  
				imprimirTexto("cadastrarNota")
				mat := capturarEntrada("matricula")
				cod := capturarEntrada("código de disciplina")
				n := capturarEntrada("nota")

			  //chamar função aqui
				//imprimirTexto(mat + cod + n)


			  	break

			case m.consultarNota:
				imprimirTexto("consultarNota")
			  
			  //chamar função aqui

			  	break

			case m.consultarNotas:
				imprimirTexto("consultarNotas")
			  
			  //chamar função aqui

			  	break

			case m.consultarCR:
				imprimirTexto("consultarCR")
			  
			  //chamar função aqui

			  	break	

			case m.terminar:
				imprimirTexto("Adeus!")
				return		  

			 default:
			 	imprimirTexto("\n*****\nERRO: OPÇÃO '" + scanner.Text() + "' INVÁLIDA!\n*****\n")
			 	m.ExibirOpcoes()

		
		}
	
	}

}



func (m *Menu) ExibirMenu() {
	imprimirTexto("\nOlá, usuário, o que deseja fazer?")
	imprimirTexto("Abaixo, segue uma ação e sua tecla correspondente.")
	imprimirTexto("Assim, caso se deseje executar a ação X, deve-se digitar a tecla Y correspondente a mesma.")
	m.configurarTeclas()
	m.ExibirOpcoes()
	m.ExecutarAcoes()



}



func main() {
	menu := Menu{}

	menu.ExibirMenu()
	
}