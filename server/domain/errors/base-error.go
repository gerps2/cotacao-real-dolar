package errors

type CustomError struct {
	Erro              error
	MensagemRetorno   string
	MensagemParaLogar string
	StatusCode        int
}
