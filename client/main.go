package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CotacaoResponse struct {
	Data struct {
		Bid string `json:"bid"`
	}
	Message string `json:"message"`
}

func main() {
	BuscarCotacao()
}

func BuscarCotacao() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Erro ao criar a requisição:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erro do servidor:", resp.Status)
		return
	}

	var cotacao CotacaoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		fmt.Println("Erro ao decodificar a resposta:", err)
		return
	}

	fmt.Printf("Cotação do dólar para real brasileiro: %s\n", cotacao.Data.Bid)
}
