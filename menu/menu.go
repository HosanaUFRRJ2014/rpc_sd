package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	. "client"

)


type Menu struct {
	Help int
	CadastrarNota int
	ConsultarNota int
	ConsultarNotas int
	ConsultarCR int
	Terminar int
}

func imprimirTexto(texto string) {
	fmt.Println(texto)
}

func (m *Menu) configurarTeclas() {
	m.Help = 0
	m.CadastrarNota = 1
	m.ConsultarNota = 2
	m.ConsultarNotas = 3
	m.ConsultarCR = 4
	m.Terminar = 9
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
	imprimirTexto("Cadastrar Nota = "+ strconv.Itoa(m.CadastrarNota))
	imprimirTexto("Consultar Nota = "+ strconv.Itoa(m.ConsultarNota))
	imprimirTexto("Consultar Notas = "+ strconv.Itoa(m.ConsultarNotas))
	imprimirTexto("Consultar CR = "+ strconv.Itoa(m.ConsultarCR))
	imprimirTexto("Exibir os atalhos novamente = "+ strconv.Itoa(m.Help))
	imprimirTexto("Terminar a execução do programa = "+ strconv.Itoa(m.Terminar))
	imprimirTexto("\n*****\n")



}


func (m Menu) ExecutarAcoes(c * Cliente) {
	
	scanner := bufio.NewScanner(os.Stdin)
	for  scanner.Scan(){

		//entrada:= scanner.Scan()
		escolha,erro := strconv.Atoi(scanner.Text())


		if erro != nil {
		  fmt.Errorf("Erro: Valor informado não é um número!")
		  fmt.Println(erro)
		}


		switch escolha {
			case m.Help:
			  	m.ExibirOpcoes()
			  	break

			case m.CadastrarNota:
			  
				imprimirTexto("CadastrarNota")
				mat := capturarEntrada("matricula")
				cod := capturarEntrada("código de disciplina")
				n := capturarEntrada("nota")

			  //chamar função aqui
				//imprimirTexto(mat + cod + n)
				nota,err := strconv.ParseFloat(n,32)
				if err != nil {
					fmt.Errorf("Erro: Nota informada não é um número!")
					fmt.Println(err)
				}
				sucesso := c.CadastrarNota(mat,cod, float32(nota))
				imprimirTexto("\nSucesso Cadastro?")
				fmt.Println(sucesso)


			  	break

			case m.ConsultarNota:
				imprimirTexto("ConsultarNota")
				mat := capturarEntrada("matricula")
				cod := capturarEntrada("código de disciplina")
			  
			  //chamar função aqui
				nota := c.ConsultarNota(mat,cod)
				imprimirTexto("\nA nota do aluno " + mat + " em " + cod + " é: ")
				fmt.Println(nota)

			  	break

			case m.ConsultarNotas:
				imprimirTexto("ConsultarNotas")
				mat := capturarEntrada("matricula")
			  
			  //chamar função aqui
				notas := c.ConsultarNotas(mat)
				imprimirTexto("\nAs notas do aluno " + mat + " são: ")
				fmt.Println(notas)

			  	break

			case m.ConsultarCR:
				imprimirTexto("ConsultarCR")
				mat := capturarEntrada("matricula")
			  
			  //chamar função aqui
				cr := c.ConsultarCR(mat)
				imprimirTexto("\nO CR do aluno " + mat + " é: ")
				fmt.Println(cr)

			  	break	

			case m.Terminar:
				imprimirTexto("Adeus!")
				return		  

			 default:
			 	imprimirTexto("\n*****\nERRO: OPÇÃO '" + scanner.Text() + "' INVÁLIDA!\n*****\n")
			 	m.ExibirOpcoes()

		
		}

		imprimirTexto("\n\n---------------\n")
		imprimirTexto("Escolha uma nova operação para ser realizada. (Tecla 0 lista todas as opções)\n")
	
	}

}



func (m *Menu) ExibirMenu(c * Cliente) {
	imprimirTexto("\nOlá, usuário, o que deseja fazer?")
	imprimirTexto("Abaixo, segue uma ação e sua tecla correspondente.")
	imprimirTexto("Assim, caso se deseje executar a ação X, deve-se digitar a tecla Y correspondente a mesma.")
	m.configurarTeclas()
	m.ExibirOpcoes()
	m.ExecutarAcoes(c)



}


