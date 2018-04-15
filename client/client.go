package client


import  (
	"fmt"
	//."github.com/HosanaUFRRJ2014/interfaceCadastroNotas" 
	."interfaceCadastroNotas"


	"net/rpc"

)

type Cliente struct {
	cliente *rpc.Client
}

func checarErro(erro error) {
	if erro!= nil {
		fmt.Println(erro)
	}
}


func (c *Cliente) CadastrarNota(matricula string, codDisciplina string, nota float32) bool{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{matricula,codDisciplina,nota}
	var sucesso bool
	erro := c.cliente.Call("CadastroNotas.CadastrarNota",dadosCadastro,&sucesso)
	checarErro(erro)

	return sucesso
}

func (c *Cliente) ConsultarNota(matricula string, codDisciplina string) float32{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{Matricula:matricula,Codigo:codDisciplina}
	var nota float32
	erro := c.cliente.Call("CadastroNotas.ConsultarNota",dadosCadastro,&nota)
	checarErro(erro)


	return nota
}


func (c *Cliente) ConsultarNotas(matricula string) []float32{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{Matricula:matricula}
	var notas []float32
	erro := c.cliente.Call("CadastroNotas.ConsultarNotas",dadosCadastro,&notas)
	checarErro(erro)


	return notas
}


func (c *Cliente) ConsultarCR(matricula string) float32{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{Matricula:matricula}
	var cr float32
	erro := c.cliente.Call("CadastroNotas.ConsultarCR",dadosCadastro,&cr)
	checarErro(erro)


	return cr
}


func ConectarComServidor() Cliente {
	

	//conexão com o servidor
	serverAddress := "127.0.0.1"

	cliente, erro := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if erro != nil {
		panic(erro)
	}

	c := Cliente{cliente:cliente}


	return c


}
