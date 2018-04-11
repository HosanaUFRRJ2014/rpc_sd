# rpc_sd 
Este repositório contém um código servidor de uma aplicação com RPC para um trabalho da disciplina de Sistemas Distribuídos.


## Configurando o ambiente de Desenvolvimento e Testes:
  Em breve, instruções de como instalar e configurar o GO.



### Importando biblioteca

Existem duas formas de importar e utilizar a biblioteca feita para esta aplicação: *localmente*, na qual é necessário obter o código fonte e *remotamente*, na qual é preciso apenas fazer com que o Go capture a referência para o repositório na Internet.

  #### Utilizando a biblioteca localmente:

  Para que a aplicação funcione, a pasta src deve estar dentro do workspace do GO. [Consulte a documentação para mais detalhes](https://golang.org/doc/install#testing). Caso seu workspace esteja em $HOME, a pasta src está em $HOME/go/. Assim, é necessário criar uma pasta "interfaceCadastroNotas" e colocar dentro dela o arquivo "interfaceCadastroNotas.go". 
  O esquema abaixo mostra como a organização dos arquivos deve ficar:
 
     src/
        interfaceCadastroNotas/
             interfaceCadastroNotas.go

        rpc_sd/
             server.go
             client.go
             README.md


  #### Utilizando a biblioteca remotamente:

  *[Em breve]*


## Para Compilar:

    go build server.go
  
    go build client.go
  
  
  
## Para Executar:
    
    ./server.go
  
  
Em outro terminal, execute:

    ./client.go
