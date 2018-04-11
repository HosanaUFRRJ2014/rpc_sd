//package server
package main


import (

	"fmt"

	"github.com/HosanaUFRRJ2014/interfaceCadastroNotas" 

	"io/ioutil"
	"os"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"


)

const CAMINHO_ARQUIVO string = "alunosCadastrados.txt"

type CadastroNotas int


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

	matriculasIguais := matricula == dadosAluno[0]

	//Se o código foi informado, este deve ser comparado. Se não, retorna true, pois está sendo usando por ConsultarNotas que não possui código como parâmetro.
	codDisciplinaIguais := true

	if codDisciplina != "" {
		codDisciplinaIguais = codDisciplina == dadosAluno[1]
	}

    return matriculasIguais && codDisciplinaIguais
}

func extrairNotaAluno(linhaArquivo string) (float32, error) {
	dadosAluno := strings.Split(linhaArquivo, "\t")
	notaAluno, erro := strconv.ParseFloat(dadosAluno[2], 32)

	return float32(notaAluno), erro
}

func copiarValoresVetor(vetorDestino *[]float32, vetorOrigem [] float32) {
	for i := 0; i < len(vetorOrigem); i++ {
		*vetorDestino = append(*vetorDestino, vetorOrigem[i])
	}
	
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

/****************************************MÉTODOS EXPORTADOS DO RPC*******************************************/

//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
//Em letra maiúscula porque os métodos precisam ser exportados
func (cadastroNotas *CadastroNotas) CadastrarNota(dadosCadastro interfaceCadastroNotas.DadosCadastro, sucessoCadastro *bool) error{
	
	var arquivo * os.File
	var sucessoModificarNota bool 
	var erro error

	//abertura de arquivo
	arquivo, erro = os.OpenFile(CAMINHO_ARQUIVO,os.O_RDWR | os.O_CREATE, 0666)
	defer arquivo.Close()


	//tentar modificar nota, caso matricula && disciplina estejam cadastradas
	sucessoModificarNota,erro = modificarNotaCadastrada(arquivo, dadosCadastro.Matricula, dadosCadastro.Codigo, dadosCadastro.Nota)


	//Senão estiverem cadastradas, adicionar nota não existente
	if !sucessoModificarNota {
		erro = salvarNovaNota(arquivo, dadosCadastro.Matricula, dadosCadastro.Codigo, dadosCadastro.Nota)		

	}

	if erro == nil {
		*sucessoCadastro = true
	}



	return erro


	
}

//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
func (cadastroNotas *CadastroNotas) ConsultarNota(dadosCadastro interfaceCadastroNotas.DadosCadastro, nota * float32) error{

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
				
		if cadastroEncontrado(linhasArquivo[i], dadosCadastro.Matricula, dadosCadastro.Codigo) {
			
			notaAluno, erro = extrairNotaAluno(linhasArquivo[i])
			*nota = float32(notaAluno)
			return erro
			
		}

		i++

	}


	erro = fmt.Errorf("Erro: Impossível encontrar nota do aluno %q na disciplina %q.", dadosCadastro.Matricula, dadosCadastro.Codigo)


	return erro

}

//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
func (cadastroNotas *CadastroNotas) ConsultarNotas(dadosCadastro interfaceCadastroNotas.DadosCadastro, notas *[] float32) error{

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
				
		if cadastroEncontrado(linhasArquivo[i], dadosCadastro.Matricula ,"") {
			notaAluno, erro = extrairNotaAluno(linhasArquivo[i])
			notasAluno = append(notasAluno, notaAluno)
			
		}

		i++

	}

	if len(notasAluno) == 0 {
		erro = fmt.Errorf("Erro: Impossível encontrar notas do aluno %q.", dadosCadastro.Matricula)	
	}


	copiarValoresVetor(notas, notasAluno)


	return erro

}

//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
func (cadastroNotas *CadastroNotas) ConsultarCR(dadosCadastro interfaceCadastroNotas.DadosCadastro, cr * float32) error {

	var notasAluno [] float32
	erro := cadastroNotas.ConsultarNotas(dadosCadastro, &notasAluno)
	*cr = -1

	if erro == nil {
		*cr = calcularCR(notasAluno)	
	} else {
		erro = fmt.Errorf("Erro: Impossível calcular CR do aluno %q.", dadosCadastro.Matricula)
	}

	return erro

}


func main() {
	cadastroNotas := new(CadastroNotas)
	rpc.Register(cadastroNotas)
	rpc.HandleHTTP()
	erro := http.ListenAndServe(":1234", nil)
    if erro != nil {
        panic(erro)
    }

   
}
