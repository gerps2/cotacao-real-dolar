package errors

import "net/http"

func NewCotacaoRepositoryCreateError(err error) CustomError {
	return CustomError{
		Erro:              err,
		MensagemRetorno:   "Erro interno, tente novamente mais tarde.",
		MensagemParaLogar: "Erro ao executar a query de inserção no banco de dados",
		StatusCode:        http.StatusInternalServerError,
	}
}
