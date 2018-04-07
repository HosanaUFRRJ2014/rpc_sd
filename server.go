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


//FIXME: Mudar struct Disciplina e Aluno para DadosCadastro

/*type Disciplina struct {
	codigo string
	nota float32
}

type Aluno struct {
	matricula string
	//disciplinas map[string]float32
	disciplinas []Disciplina
}*/


type DadosCadastro struct{
	matricula string
	codigo string
	nota float32
}

type CadastroNotas struct {
	alunos []DadosCadastro
}

//type CadastroNotas int

func checarErro(erro error) {
    if erro != nil {
       fmt.Print(erro, "\n")
       panic(erro)
     //  os.Exit(1)
    }
}

func construirEscritaArquivo(matricula string, codDisciplina string, nota float32) []byte{
	byteSlice := []byte(matricula + "\t" + codDisciplina + "\t" + strconv.FormatFloat(float64(nota),'E', -1, 64) + "\n")
	return byteSlice

}

func moverPonteiroArquivo(arquivo *os.File, offset int64, origem int) (int64,error){
	posAtual, erro := arquivo.Seek(offset, origem) 
	return posAtual, erro
	
}

func recuperarCadastros(arquivo *os.File) ([]string, int64, error) {
	dadosArquivo, erro := ioutil.ReadAll(arquivo)
	textoArquivo := string(dadosArquivo)
	linhasArquivo := strings.Split(textoArquivo,"\n") 
	tamanhoArquivo := int64(len(dadosArquivo))

	return linhasArquivo, tamanhoArquivo, erro
}

func cadastroEncontrado(linhaArquivo string, matricula string, codDisciplina string) bool{
	
	if len(linhaArquivo) == 0 {
		return false
	}

	var dadosAluno  [] string
	dadosAluno = strings.Split(linhaArquivo,"\t")

	if matricula == dadosAluno[0] && codDisciplina == dadosAluno[1] {
		return true
	}

	return false
}

func modificarNotaCadastrada(arquivo * os.File, matricula string, codDisciplina string, nota float32) (bool, error){

//	var linhasArquivo [] string
	var tamanhoLinhaAtual int64
	var i int
	var posAtual int64
//	var erro error
/*	dadosArquivo, erro := ioutil.ReadAll(arquivo)
	


	textoArquivo := string(dadosArquivo)
	linhasArquivo = strings.Split(textoArquivo,"\n") */

	linhasArquivo,tamanhoArquivo, erro := recuperarCadastros(arquivo)


	if tamanhoArquivo == 0 {
		return false,nil
	}


	//colocar ponteiro no início do arquivo
	posAtual,erro = moverPonteiroArquivo(arquivo, 0, os.SEEK_SET) 


	for i < len(linhasArquivo) {
				
		tamanhoLinhaAtual = int64(len(linhasArquivo[i]))
		
		if cadastroEncontrado(linhasArquivo[i], matricula ,codDisciplina) {
			bytesAEscrever := construirEscritaArquivo(matricula,codDisciplina,nota)
			arquivo.WriteAt(bytesAEscrever, posAtual)
			return true,nil
			
		}

		//mover ponteiro para a próxima linha
		posAtual,erro = moverPonteiroArquivo(arquivo,tamanhoLinhaAtual + 1, os.SEEK_CUR)
		i++


	}

	//mover ponteiro para o fim do arquivo
	posAtual,erro = moverPonteiroArquivo(arquivo,tamanhoArquivo, os.SEEK_SET)
	return false,erro


	
}

func cadastrarNovaNota(arquivo *os.File, matricula string, codDisciplina string, nota float32) error{
	escrita := construirEscritaArquivo(matricula, codDisciplina, nota)
	_, erro := arquivo.Write(escrita)
	return erro
}





func salvar(matricula string, codDisciplina string, nota float32) error {

	var arquivo * os.File
	var sucessoModificarNota bool 
	var erro error

	//abertura de arquivo
	arquivo, erro = os.OpenFile(CAMINHO_ARQUIVO,os.O_RDWR | os.O_CREATE, 0666)
	defer arquivo.Close()


	//tentar modificar nota, caso matricula && disciplina estejam cadastradas
	sucessoModificarNota,erro = modificarNotaCadastrada(arquivo, matricula, codDisciplina, nota)


	//Senão estiverem cadastradas, adicionar nota não existente
	if !sucessoModificarNota {
		erro = cadastrarNovaNota(arquivo, matricula, codDisciplina, nota)		

	}


	return erro
	
}


func consultarNota(matricula string, codDisciplina string) (float32, error) {
	
	var arquivo * os.File
	var notaAluno float64
	var erro error
	var linhasArquivo []string
	var i int


	//abertura de arquivo
	arquivo, erro = os.OpenFile(CAMINHO_ARQUIVO,os.O_RDWR | os.O_CREATE, 0666)
	defer arquivo.Close()


	linhasArquivo, _ , erro = recuperarCadastros(arquivo)

	for i < len(linhasArquivo) {
				
		if cadastroEncontrado(linhasArquivo[i], matricula ,codDisciplina) {
			dadosAluno := strings.Split(linhasArquivo[i], "\t")
			notaAluno, erro = strconv.ParseFloat(dadosAluno[2], 32)
			return float32(notaAluno), erro
			
		}

		i++

	}


	erro = fmt.Errorf("Erro: Impossível encontrar nota do aluno %q na disciplina %q.", matricula, codDisciplina)

	return float32(notaAluno), erro



	
}


//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
func (cadastroNotas *CadastroNotas) cadastrarNota(dadosCadastro DadosCadastro, sucessoCadastro *bool) error{
	
	erro :=	salvar(dadosCadastro.matricula, dadosCadastro.codigo, dadosCadastro.nota)

	if erro == nil {
		*sucessoCadastro = true
	}

	return erro

	
}

//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
func (cadastroNotas *CadastroNotas) consultarNota(dadosCadastro DadosCadastro, nota * float32) error{
	var erro error
	*nota, erro = consultarNota(dadosCadastro.matricula, dadosCadastro.codigo)
	return erro

}

//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
/*func (cadastroNotas *CadastroNotas) consultarNotas(dadosCadastro DadosCadastro, notas *[] float32) error{
	var erro error
	*notas

}*/


func main() {

	c := CadastroNotas{}

	//FIXME : modificar parametros para não dar erro
	/*erro := c.cadastrarNota("2014780267",&Disciplina{"IM888",1.0})
	checarErro(erro) //client deve receber o erro
	c.cadastrarNota("2014780267",&Disciplina{"IM556",2.0})
	c.cadastrarNota("2014780267",&Disciplina{"AA944",3.0})
	c.cadastrarNota("2014780267",&Disciplina{"BB331",4.0})
	c.cadastrarNota("2014780267",&Disciplina{"IM556",5.0})*/


	var notaAluno float32

	erro := c.consultarNota(DadosCadastro{matricula:"2014780267",codigo:"IA556"}, &notaAluno)

	checarErro(erro)

	fmt.Println(notaAluno)


	fmt.Println("")
	
//	fmt.Println(c.alunos)

/*	a := Aluno{matricula: "2014780267"}

	a.adicionarDisciplinaNota("IM887", 10.0)
	a.adicionarDisciplinaNota("AA888", 8.0)
	a.adicionarDisciplinaNota("IM887", 5.0)

	fmt.Println(a.disciplinas)*/
}