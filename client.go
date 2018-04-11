package main


import  (
	"fmt"
	."github.com/HosanaUFRRJ2014/interfaceCadastroNotas" 

	"net/rpc"

)

type _CadastroNotas struct {
	cliente *rpc.Client
}

func checarErro(erro error) {
	if erro!= nil {
		fmt.Println(erro)
	}
}


func (c *_CadastroNotas) cadastrarNota(matricula string, codDisciplina string, nota float32) bool{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{matricula,codDisciplina,nota}
	var sucesso bool
	erro := c.cliente.Call("CadastroNotas.CadastrarNota",dadosCadastro,&sucesso)
	checarErro(erro)

	return sucesso
}

func (c *_CadastroNotas) consultarNota(matricula string, codDisciplina string) float32{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{Matricula:matricula,Codigo:codDisciplina}
	var nota float32
	erro := c.cliente.Call("CadastroNotas.ConsultarNota",dadosCadastro,&nota)
	checarErro(erro)


	return nota
}


func (c *_CadastroNotas) consultarNotas(matricula string) []float32{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{Matricula:matricula}
	var notas []float32
	erro := c.cliente.Call("CadastroNotas.ConsultarNotas",dadosCadastro,&notas)
	checarErro(erro)


	return notas
}


func (c *_CadastroNotas) consultarCR(matricula string) float32{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{Matricula:matricula}
	var cr float32
	erro := c.cliente.Call("CadastroNotas.ConsultarCR",dadosCadastro,&cr)
	checarErro(erro)


	return cr
}


func main() {
	
	 /*if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], "server")
        os.Exit(1)
    }*/

	//conexão com o servidor
	serverAddress := "127.0.0.1"

	cliente, erro := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if erro != nil {
		panic(erro)
	}

	c := _CadastroNotas{cliente:cliente}

	var sucesso bool

	sucesso = c.cadastrarNota("2014780267", "AA5879", 8.5)
	fmt.Println(sucesso)


	sucesso = c.cadastrarNota("2014780267", "BB5569", 7.0)
	fmt.Println(sucesso)


	nota := c.consultarNota("2013478522", "DD5689")
	fmt.Println(nota)

	notas := c.consultarNotas("2014780267") //deveria lançar erro, mas não lança
	fmt.Println(notas)

	cr := c.consultarCR("2014780267")
	fmt.Println(cr)


	//Organizar melhor métodos do client, procurando evitar repetição de código.


	//export GOPATH=$HOME/go
}
