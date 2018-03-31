//package server
package main

import (

	"fmt"
/*	"net/rpc"
	"errors"
*/
)


type Disciplina struct {
	codigo string
	nota float32
}


//var disciplinaNota map[string]Disciplina

type Aluno struct {
//	matricula string
	disciplinas map[string]float32

}

type CadastroNotas struct {
	alunos map[string]Aluno
}

/*
func (aluno *Aluno) criarMapaDisciplinas(alu) {
	
}*/


/*func (aluno *Aluno) adicionarDisciplinaNota( codDisciplina string, nota float32) {
	
	if aluno.disciplinas == nil {
		aluno.disciplinas = make(map[string]float32)
	}

	aluno.disciplinas[codDisciplina] = nota


}*/


//FIXME: OS pararâmetros não são os mesmos dos pedidos na descrição do exercício!
func (cadastroNotas *CadastroNotas) cadastrarNota(matricula string, disciplina *Disciplina) error{
	
	if cadastroNotas.alunos == nil {
		//cadastroNotas.aluno := Aluno{}
		cadastroNotas.alunos = make(map[string]Aluno)
		//cadastroNotas.alunos[matricula].disciplinas = make(map[string]float32)
	}

	aluno := Aluno{}
	aluno.disciplinas = make(map[string]float32)
	aluno.disciplinas[disciplina.codigo] = disciplina.nota
	cadastroNotas.alunos[matricula] = aluno
	//cadastroNotas.alunos.disciplinas[disciplina.codigo] = disciplina.nota


	//testando se o elemento foi adicionado no mapa com sucesso
	/*if aluno.disciplinas[disciplina.codigo] != nil
		*resposta = true*/

	return nil

	
}




/*type CadastroNotas int


func (tipo *CadastroNotas) cadastrarNota (aluno map[string], resposta *bool) {

	//checar se nota existe. 
		//Se existe, subscrever
		//Caso contrário, criar novo cadastro no final do arquivo




	*resposta = false
	return nil
}

*/

func main() {

	c := CadastroNotas{}

	c.cadastrarNota("2014780267",&Disciplina{"IM887",6.0})
	c.cadastrarNota("2014780267",&Disciplina{"IM888",7.0})

	
	fmt.Println(c.alunos)

/*	a := Aluno{matricula: "2014780267"}

	a.adicionarDisciplinaNota("IM887", 10.0)
	a.adicionarDisciplinaNota("AA888", 8.0)
	a.adicionarDisciplinaNota("IM887", 5.0)

	fmt.Println(a.disciplinas)*/
}