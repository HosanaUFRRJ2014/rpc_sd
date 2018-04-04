//package server
package main

import (

//	"bufio"
	"fmt"
//	"log"
	"io/ioutil"
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

func checarErro(erro error) {
    if erro != nil {
        panic(erro)
    }
}

func (aluno *Aluno) construirEscritaArquivo(codDisciplina string, nota float32) []byte{
	byteSlice := []byte(aluno.matricula + "\t" + codDisciplina + "\t" + strconv.FormatFloat(float64(nota),'E', -1, 64) + "\n")
	return byteSlice

}


func (aluno *Aluno) cadastroEncontrado(linhaArquivo string, codDisciplina string) bool{
	
	if len(linhaArquivo) == 0 {
		return false
	}

	var dadosAluno  [] string
	dadosAluno = strings.Split(linhaArquivo,"\t")
  //  fmt.Println(linhaArquivo, dadosAluno[1])

	fmt.Println("matricula: ", aluno.matricula, " && ", dadosAluno[0])
	fmt.Println("codDisciplina: ", codDisciplina, " && ", dadosAluno[1])
	fmt.Println("")

	if aluno.matricula == dadosAluno[0] && codDisciplina == dadosAluno[1] {
		return true
	}

	return false
}

func moverPonteiroArquivo(arquivo *os.File, offset int64, origem int) {
	movidos, erro := arquivo.Seek(offset, origem) 
	fmt.Println("Movidos: ", movidos)
	checarErro(erro)
	
}

//Provável FIXME: arquivo não ser do tipo File, necessidade de descobrir o tipo correto
func (aluno *Aluno) modificarNotaCadastrada(arquivo * os.File, codDisciplina string, nota float32) bool{

	var linhasArquivo [] string
	var tamanhoLinhaAtual int64
	var i int
	dadosArquivo, erro := ioutil.ReadAll(arquivo)
	checarErro(erro)


	textoArquivo := string(dadosArquivo)
	linhasArquivo = strings.Split(textoArquivo,"\n") 
//	fmt.Println("Conteudo:\n\n ", textoArquivo)

	if len(textoArquivo) == 0 {
		return false
	}


	//colocar ponteiro no início do arquivo
	moverPonteiroArquivo(arquivo, 0, os.SEEK_SET) 


	fmt.Println("Len:   ", len(linhasArquivo))


	for i < len(linhasArquivo) {
		fmt.Println("\nlinhaArquivo:   ",len(linhasArquivo[i]))
		tamanhoLinhaAtual += int64(len(linhasArquivo[i]))
		if aluno.cadastroEncontrado(linhasArquivo[i],codDisciplina) {
			bytesAEscrever := aluno.construirEscritaArquivo(codDisciplina,nota)
			arquivo.WriteAt(bytesAEscrever, tamanhoLinhaAtual)
			return true
			
		}

		//mover ponteiro para a próxima linha
		moverPonteiroArquivo(arquivo,int64(tamanhoLinhaAtual), os.SEEK_CUR)
		//i = tamanhoLinhaAtual
		i++
		fmt.Println("i: ", i)
		//fmt.Println("Executou!!!")


	}

	//mover ponteiro para o fim do arquivo
	moverPonteiroArquivo(arquivo,int64(len(dadosArquivo)), os.SEEK_SET)
	return false


	
}

func (aluno *Aluno) salvar(codDisciplina string, nota float32) error {

	var arquivo * os.File

	arquivo, erro := os.OpenFile(CAMINHO_ARQUIVO,os.O_RDWR|/*os.O_APPEND|*/os.O_CREATE,0666)
	checarErro(erro)
	defer arquivo.Close()

	sucessoModificarNota := aluno.modificarNotaCadastrada(arquivo, codDisciplina, nota)

	fmt.Println("Sucesso modificar  nota?", sucessoModificarNota)

	if !sucessoModificarNota {
		//adicionar nota não existente
		// Provável FIXME: Colocar cursor para o fim do arquivo antes da escrita
		escrita := aluno.construirEscritaArquivo(codDisciplina, nota)
   		_, erro := arquivo.Write(escrita)
   		checarErro(erro)

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

	c.cadastrarNota("2014780267",&Disciplina{"IM556",9.0})
	c.cadastrarNota("2014780267",&Disciplina{"IM888",7.0})
	c.cadastrarNota("2014780267",&Disciplina{"AA944",5.0})
	c.cadastrarNota("2014780267",&Disciplina{"IM556",10.0})

	
//	fmt.Println(c.alunos)

/*	a := Aluno{matricula: "2014780267"}

	a.adicionarDisciplinaNota("IM887", 10.0)
	a.adicionarDisciplinaNota("AA888", 8.0)
	a.adicionarDisciplinaNota("IM887", 5.0)

	fmt.Println(a.disciplinas)*/
}