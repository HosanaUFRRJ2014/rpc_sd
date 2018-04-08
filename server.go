//package server
package server

import (

//	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/rpc"
	"regexp"
	"strconv"
	"strings"

/*	"errors" */

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
      // panic(erro)
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

	matriculasIguais, _ := regexp.MatchString(matricula, dadosAluno[0])
	codDisciplinaIguais, _ := regexp.MatchString(codDisciplina, dadosAluno[1])

    return matriculasIguais && codDisciplinaIguais
}

func extrairNotaAluno(linhaArquivo string) (float32, error) {
	dadosAluno := strings.Split(linhaArquivo, "\t")
	notaAluno, erro := strconv.ParseFloat(dadosAluno[2], 32)

	return float32(notaAluno), erro
}

func modificarNotaCadastrada(arquivo * os.File, matricula string, codDisciplina string, nota float32) (bool, error){
	var tamanhoLinhaAtual int64
	var i int
	var posAtual int64

	linhasArquivo,tamanhoArquivo, erro := recuperarCadastros(arquivo)


	if tamanhoArquivo == 0 {
		return false,erro
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

func salvarNovaNota(arquivo *os.File, matricula string, codDisciplina string, nota float32) error{
	escrita := construirEscritaArquivo(matricula, codDisciplina, nota)
	_, erro := arquivo.Write(escrita)
	return erro
}


func  calcularCR(notasAluno [] float32) (float32){
	var cr float32
	numDisciplinasAluno := len(notasAluno)

	for i := 0; i < numDisciplinasAluno ; i++ {
		cr += notasAluno[i]	
	}

	return cr/float32(numDisciplinasAluno)
}

func cadastrarNota(matricula string, codDisciplina string, nota float32) error {

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
		erro = salvarNovaNota(arquivo, matricula, codDisciplina, nota)		

	}


	return erro
	
}


func consultarNota(matricula string, codDisciplina string) (float32, error) {
	
	var arquivo * os.File
	var notaAluno float32
	var erro error
	var linhasArquivo []string
	var i int


	//abertura de arquivo
	arquivo, erro = os.OpenFile(CAMINHO_ARQUIVO,os.O_RDWR | os.O_CREATE, 0666)
	defer arquivo.Close()


	linhasArquivo, _ , erro = recuperarCadastros(arquivo)

	for i < len(linhasArquivo) {
				
		if cadastroEncontrado(linhasArquivo[i], matricula ,codDisciplina) {
			
			notaAluno, erro = extrairNotaAluno(linhasArquivo[i])
			return notaAluno, erro
			
		}

		i++

	}


	erro = fmt.Errorf("Erro: Impossível encontrar nota do aluno %q na disciplina %q.", matricula, codDisciplina)

	return float32(notaAluno), erro



	
}

func consultarNotas(matricula string) ([] float32, error){
	var arquivo * os.File
	var notaAluno float32
	var erro error
	var linhasArquivo []string
	var i int
	var notasAluno [] float32


	//abertura de arquivo
	arquivo, erro = os.OpenFile(CAMINHO_ARQUIVO,os.O_RDWR | os.O_CREATE, 0666)
	defer arquivo.Close()


	linhasArquivo, _ , erro = recuperarCadastros(arquivo)

	for i < len(linhasArquivo) {
				
		if cadastroEncontrado(linhasArquivo[i], matricula ,".") {
			notaAluno, erro = extrairNotaAluno(linhasArquivo[i])
			notasAluno = append(notasAluno, notaAluno)
			
		}

		i++

	}

	if len(notasAluno) == 0 {
		erro = fmt.Errorf("Erro: Impossível encontrar notas do aluno %q.", matricula)	
	}

	return notasAluno, erro

}

func consultarCR(matricula string) (float32, error){
	var cr float32
	notasAluno, erro := consultarNotas(matricula)
	cr = -1

	if erro == nil {
		cr = calcularCR(notasAluno)	
	}

	/*else {
		erro = fmt.Errorf("Erro")
	}*/

	return cr, erro
}


//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
func (cadastroNotas *CadastroNotas) cadastrarNota(dadosCadastro DadosCadastro, sucessoCadastro *bool) error{
	
	erro :=	cadastrarNota(dadosCadastro.matricula, dadosCadastro.codigo, dadosCadastro.nota)

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
func (cadastroNotas *CadastroNotas) consultarNotas(dadosCadastro DadosCadastro, notas *[] float32) error{
	var erro error
	*notas,erro = consultarNotas(dadosCadastro.matricula)
	return erro

}

//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
func (cadastroNotas *CadastroNotas) consultarCR(dadosCadastro DadosCadastro, cr * float32) error {
	var erro error

	*cr, erro = consultarCR(dadosCadastro.matricula)
	return erro
}



cadastroNotas := New(CadastroNotas)
rpc.Register(cadastroNotas)
rpc.HadleHttp(cadastroNotas)
listenner, erro := net.Listen("tcp", ":3000")
checarErro(erro)

go http.Serve(listenner, nil)