package main


import  (
	"fmt"
	. "interfaceCadastroNotas"
	"net/rpc"

)

type _CadastroNotas struct {
	cliente *rpc.Client
}


func (c *_CadastroNotas) cadastrarNota(matricula string, codDisciplina string, nota float32) bool{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{matricula,codDisciplina,nota}
	var sucesso bool
	erro := c.cliente.Call("CadastroNotas.CadastrarNota",dadosCadastro,&sucesso)
	if erro!= nil {
		panic(erro)
	}

	return sucesso
}

func (c *_CadastroNotas) consultarNota(matricula string, codDisciplina string) float32{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{Matricula:matricula,Codigo:codDisciplina}
	var nota float32
	erro := c.cliente.Call("CadastroNotas.ConsultarNota",dadosCadastro,&nota)
	if erro!= nil {
		panic(erro)
	}

	return nota
}


func (c *_CadastroNotas) consultarNotas(matricula string) []float32{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{Matricula:matricula}
	var notas []float32
	erro := c.cliente.Call("CadastroNotas.ConsultarNotas",dadosCadastro,&notas)
	if erro!= nil {
		panic(erro)
	}

	return notas
}


func (c *_CadastroNotas) consultarCR(matricula string) float32{
	//realizando chamada síncrona
	dadosCadastro := &DadosCadastro{Matricula:matricula}
	var cr float32
	erro := c.cliente.Call("CadastroNotas.ConsultarCR",dadosCadastro,&cr)
	if erro!= nil {
		panic(erro)
	}

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

	//realizando chamada síncrona
/*	dadosCadastro := DadosCadastro{"2014780267","CC5555",9.54}
	var sucesso bool
	erro = cliente.Call("CadastroNotas.CadastrarNota",dadosCadastro,&sucesso)
	if erro!= nil {
		panic(erro)
	}
	fmt.Println(sucesso)*/

	c := _CadastroNotas{cliente:cliente}

	sucesso := c.cadastrarNota("2013478522", "DD5689", 9.5)
	fmt.Println(sucesso)

	nota := c.consultarNota("2013478522", "DD5689")
	fmt.Println(nota)

	notas := c.consultarNotas("20134") //deveria lançar erro, mas não lança
	fmt.Println(notas)

	cr := c.consultarCR("2013478522")
	fmt.Println(cr)


	//FIXME: dando match em prefixo de matrícula e disciplina, ao invés de palavra exata
	//TODO: Remover panic e colocar os.Exit(1)
	//TODO: desencapsular métodos receptors do server
	//Organizar melhor métodos do client, procurando evitar repetição de código.
}