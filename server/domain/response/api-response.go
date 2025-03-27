package response

import (
	"encoding/json"
	"log"
	"net/http"
	"server/domain/errors"
)

type ApiResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Responder(w http.ResponseWriter, data interface{}, err *errors.CustomError) {
	response := ApiResponse{}

	if err != nil {
		log.Printf("Erro: %v, Stack Trace: %+v", err.MensagemParaLogar, err.Erro)
		response.Message = err.MensagemRetorno
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.StatusCode)
	} else {
		response.Data = data
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
