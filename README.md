# rpc_sd 
Este repositório contém um código servidor de uma aplicação com RPC para um trabalho da disciplina de Sistemas Distribuídos.


## Descrição Geral:

  O repositório consiste nos arquivos listados abaixo:

    client/     
         client.go
    
    interfaceCadastroNotas/
         interfaceCadastroNotas.go

    main/
         main.go

    menu/
        menu.go     
    
    server/
         server.go
    
    .gitignore
    README.md


O pacote *"server"* pertence ao lado do servidor e os pacotes *"client"*, *"menu"* e *"main"* pertencem ao lado do cliente. O pacote *"interfaceCadastroNotas"* descreve uma interface que deve ser implementada por ambos os lados. Isso foi feito a fim de se evitar repetição de código.



## Configurando o ambiente de Desenvolvimento e Testes:
  *[Em breve, instruções de como instalar e configurar o GO.]*

  Caso já o tenha instalado, pode ser que em um computador Linux Ubuntu seja necessário executar ```export GOPATH=$PATH:$HOME/go ``` toda vez que se abrir um novo terminal. 



### Utilizando as bibliotecas locais:

  Para que a aplicação funcione, as pastas do repositório devem estar dentro de go/src que, por sua vez, deve ser o workspace do GO. [Consulte a documentação para mais detalhes](https://golang.org/doc/install#testing)
 A organização dos arquivos deve estar igual a ilustrada na seção ["Descrição Geral"](## Descrição Geral). (FIX link)


## Para Compilar:

  Na pasta "src", digite os seguintes comandos:  
```
go build server/server.go
```
```  
go build client/client.go
```
```
go build menu/menu.go
```
```
go build main/main.go
```

  **IMPORTANTE:** É de **EXTREMA** importância respeitar a ordem de compilação informada acima, para que não hajam problemas de utlização de versões obsoletas de bibliotecas.

  Um fato interessante a se mencionar é que apenas os arquivos "server.go" "main.go" terão executáveis.
  
  
  
## Para Executar:
   Em um terminal, entre na pasta *"server"* e execute:
   
    ./server.go
  
  
   Em outro terminal, entre na pasta *"main"* e execute:

    ./client.go
