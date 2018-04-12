package interfaceCadastroNotas

type DadosCadastro struct {
	Matricula string
	Codigo string
	Nota float32
}


type CadastroNotas interface {
	CadastrarNota(dadosCadastro DadosCadastro, sucessoCadastro *bool) error
	ConsultarNota(dadosCadastro DadosCadastro, nota * float32) error
	ConsultarNotas(dadosCadastro DadosCadastro, notas *[] float32) error
	ConsultarCR(dadosCadastro DadosCadastro, cr * float32) error
}



