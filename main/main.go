package main

import (
	 "client"
	."menu"

)


func main() {

	c := client.ConectarComServidor()
	menu := Menu{}

	/*var sucesso bool

	sucesso = c.CadastrarNota("2014780267", "AA5879", 8.5)
	fmt.Println(sucesso)
*/
	menu.ExibirMenu(&c)
	
}
