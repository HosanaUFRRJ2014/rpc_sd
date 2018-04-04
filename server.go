//package server
package main

import (

	"bufio"
	"fmt"
//	"log"
	"os"
	"strconv"
	"strings"
/*	"net/rpc"
	"errors" */

)

const CAMINHO_ARQUIVO string = "alunosCadastrados.txt"

type Disciplina struct {
	codigo string
	nota float32
}


//var disciplinaNota map[string]Disciplina

type Aluno struct {
	matricula string
	disciplinas map[string]float32
}

type CadastroNotas struct {
	alunos []Aluno
}

func (aluno *Aluno) construirEscritaArquivo(codDisciplina string, nota float32) string{
	return aluno.matricula + "\t" + codDisciplina + "\t" + strconv.FormatFloat(float64(nota),'E', -1, 64) // + "\n"
}


func (aluno *Aluno) cadastroEncontrado(linhaArquivo string, codDisciplina string) bool{
	
	var dadosAluno  [] string
	dadosAluno = strings.Split(linhaArquivo,"\t")

  //  fmt.Println(linhaArquivo, dadosAluno[1])
	if codDisciplina == dadosAluno[1] {
		return true
	}

	return false
}

//Provável FIXME: arquivo não ser do tipo File, necessidade de descobrir o tipo correto
func (aluno *Aluno) modificarNotaCadastrada(arquivo * os.File, codDisciplina string, nota float32) bool{
	escritor := bufio.NewWriter(arquivo)
	leitor := bufio.NewScanner(arquivo)

	for leitor.Scan() {
		if aluno.cadastroEncontrado(leitor.Text(), codDisciplina) {
			// Provável FIXME: escritor não vai saber a posição do arquivo do leitor
			fmt.Fprintln(escritor, aluno.construirEscritaArquivo(codDisciplina, nota))	
		//	fmt.Println(aluno.construirEscritaArquivo(codDisciplina,nota))
			return	true	
		}
		
	}

	return false
	
}

func (aluno *Aluno) salvar(codDisciplina string, nota float32) error {

	var arquivo * os.File

	arquivo, erro := os.Create(CAMINHO_ARQUIVO)

	if erro != nil {
	//	arquivo, erro := os.Create(CAMINHO_ARQUIVO)
	//	fmt.Println("label1")

		//_ = arquivo

		/*if erro != nil {
			fmt.Println("label2")
			return erro
			
		}*/

		return erro
	}



	defer arquivo.Close()

	leitor := bufio.NewReader(arquivo)
	escritor := bufio.NewWriter(arquivo)
	leitorEscritor := bufio.NewReadWriter(leitor, escritor)	
	sucessoModificarNota := aluno.modificarNotaCadastrada(arquivo, codDisciplina, nota)

	if !sucessoModificarNota {
		//adicionar nota não existente
		// Provável FIXME: Colocar cursor para o fim do arquivo antes da escrita
		escrita := aluno.construirEscritaArquivo(codDisciplina, nota)
	//	fmt.Println(escrita)
		fmt.Fprintln(leitorEscritor, escrita)
		return leitorEscritor.Flush()

	}


	return nil
	
}



//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
func (cadastroNotas *CadastroNotas) cadastrarNota(matricula string, disciplina *Disciplina) error{
	
	var aluno Aluno	
	aluno = Aluno{matricula: matricula}
	//aluno.disciplinas[disciplina.codigo] = disciplina.nota
	aluno.salvar(disciplina.codigo, disciplina.nota)


	return nil

	
}




func main() {

	c := CadastroNotas{}

	c.cadastrarNota("2014780267",&Disciplina{"IM887",9.0})
	c.cadastrarNota("2014780267",&Disciplina{"IM888",7.0})

	
	fmt.Println(c.alunos)

/*	a := Aluno{matricula: "2014780267"}

	a.adicionarDisciplinaNota("IM887", 10.0)
	a.adicionarDisciplinaNota("AA888", 8.0)
	a.adicionarDisciplinaNota("IM887", 5.0)

	fmt.Println(a.disciplinas)*/
}