package errors

import "net/http"

func NewDatabaseConnetionError(err error) CustomError {
	return CustomError{
		Erro:              err,
		MensagemRetorno:   "Erro interno, tente novamente mais tarde.",
		MensagemParaLogar: "Erro ao abrir a conexão com o banco de dados",
		StatusCode:        http.StatusInternalServerError,
	}
}

func NewDatabaseCreateTableError(err error) CustomError {
	return CustomError{
		Erro:              err,
		MensagemRetorno:   "Erro interno, tente novamente mais tarde.",
		MensagemParaLogar: "Erro ao tentar criar a tabela no banco de dados",
		StatusCode:        http.StatusInternalServerError,
	}
}
