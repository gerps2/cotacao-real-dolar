package main

import (
	"os"
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
	bid, err := buscarCotacao()
	if err != nil {
		fmt.Println("Erro ao buscar cotação:", err)
		return
	}

	fmt.Printf("Cotação do dólar para real brasileiro: %s\n", bid)

	err = salvarCotacaoNoArquivo(bid)
	if err != nil {
		fmt.Println("Erro ao salvar a cotação no arquivo:", err)
		return
	}

	fmt.Println("Cotação salva no arquivo cotacao.txt com sucesso.")
}

func buscarCotacao() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return "", fmt.Errorf("erro ao criar a requisição: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erro do servidor: %s", resp.Status)
	}

	var cotacao CotacaoResponse
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		return "", fmt.Errorf("erro ao decodificar a resposta: %w", err)
	}

	return cotacao.Data.Bid, nil
}


func salvarCotacaoNoArquivo(bid string) error {
	conteudo := fmt.Sprintf("Dólar: %s", bid)
	err := os.WriteFile("cotacao.txt", []byte(conteudo), 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %w", err)
	}
	return nil
}

