package main


import  (
	"fmt"
	. "interfaceCadastroNotas"

	"net/rpc"




)



func main() {
	
	//conexão com o servidor
	 /*if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], "server")
        os.Exit(1)
    }*/

	serverAddress := "127.0.0.1"

	cliente, erro := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if erro != nil {
		panic(erro)
	}

	//realizando chamada síncrona
	dadosCadastro := DadosCadastro{"2014780267","CC5555",9.54}
	var sucesso bool
	erro = cliente.Call("CadastroNotas.CadastrarNota",dadosCadastro,&sucesso)
	if erro!= nil {
		panic(erro)
	}
	fmt.Println(sucesso)


}