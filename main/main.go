package main

import (
	 "client"
	."menu"

)


func main() {

	c := client.ConectarComServidor()
	menu := Menu{}

	menu.ExibirMenu(&c)
	
}
