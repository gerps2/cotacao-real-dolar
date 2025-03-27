package errors

import "net/http"

func NewTimeoutError(err error) CustomError {
	return CustomError{
		Erro:              err,
		MensagemRetorno:   "O servidor demorou muito para responder. Tente novamente mais tarde.",
		MensagemParaLogar: "Timeout ao chamar o endpoint",
		StatusCode:        http.StatusGatewayTimeout,
	}
}

func NewStatusCodeError(err error) CustomError {
	return CustomError{
		Erro:              err,
		MensagemRetorno:   "Erro ao chamar o serviço externo.",
		MensagemParaLogar: "Status code diferente de 200 ao chamar o endpoint",
		StatusCode:        http.StatusBadGateway,
	}
}

func NewUnmarshalError(err error) CustomError {
	return CustomError{
		Erro:              err,
		MensagemRetorno:   "Erro ao processar a resposta do serviço externo.",
		MensagemParaLogar: "Erro ao deserializar o JSON da resposta",
		StatusCode:        http.StatusInternalServerError,
	}
}

func NewRequestError(err error) CustomError {
	return CustomError{
		Erro:              err,
		MensagemRetorno:   "Erro ao fazer a requisição.",
		MensagemParaLogar: "Erro ao fazer a requisição para o endpoint",
		StatusCode:        http.StatusInternalServerError,
	}
}
