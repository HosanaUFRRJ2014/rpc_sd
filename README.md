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


O pacote *"server"* pertence ao lado do servidor e os pacotes *"client"*, *"menu"* e *"main"* pertencem ao lado do cliente. O pacote *"interfaceCadastroNotas"* descreve uma interface que deve ser implementada por ambos os lados.

- O pacote *"server"* contém funções de criação e gerenciamento do servidor;
- O pacote *"client"* contém funções de criação, gerenciamento do cliente e chamada do servidor;
- O pacote *"interfaceCadastroNotas"* é uma interface que deve ser implementada pelo servidor e pelo cliente, a fim de evitar repetição de código;
- O pacote *"menu"* contém funções para a exibição do menu da aplicação;
- O pacote *"main"* é o pacote principal do cliente, no qual o programa é executado.



## Configurando o ambiente de Desenvolvimento e Testes:

  Consulte o [tutorial oficial](https://golang.org/doc/install) da linguagem Go para a instalação.

  Caso já tenha instalado o Go, pode ser que em um computador Linux Ubuntu seja necessário executar ```export GOPATH=$PATH:$HOME/go ``` toda vez que se abrir um novo terminal. 



### Utilizando as bibliotecas locais:

  Para que a aplicação funcione, as pastas do repositório devem estar dentro de go/src que, por sua vez, deve ser o workspace do GO. [Consulte a documentação para mais detalhes](https://golang.org/doc/install#testing). 
 A organização dos arquivos deve estar igual a ilustrada na seção ["Descrição Geral"](https://github.com/HosanaUFRRJ2014/rpc_sd/blob/rpc/README.md#descri%C3%A7%C3%A3o-geral).


## Para Compilar:

  Na pasta "src", digite os seguintes comandos:  
```
cd server
go build server.go
cd ..
```
```  
cd client
go build client.go
cd ..
```
```
cd menu
go build menu.go
cd ..
```
```
cd main
go build main.go
cd ..
```

  **NOTA:** É de **EXTREMA** importância respeitar a ordem de compilação informada acima, para que não hajam problemas de utlização de versões obsoletas de bibliotecas.

  Um fato interessante a se mencionar é que apenas os arquivos "server.go" "main.go" terão executáveis.
  
  
  
## Para Executar:
   Em um terminal, entre na pasta *"server"* (`cd server`) e execute:
   
    ./server.go
  
  
   Em outro terminal, entre na pasta *"main"* (```cd main```) e execute:

    ./main.go


Obs: O arquivo txt que (*alunosCadastrados.txt*) salva as informações cadastradas na aplicação será criado dentro da pasta *"server"*.
