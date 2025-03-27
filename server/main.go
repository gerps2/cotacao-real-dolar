package main

import (
	"context"
	"fmt"
	"net/http"
	"server/domain/response"
	"server/repositorys"
	rest_clients "server/rest-clients"
	"time"
)

func main() {
	db, err := repositorys.InitDB("cotacoes.db")

	if err != nil {
		panic(err.Erro)
	}

	cotacaoRepo := repositorys.NewCotacaoRepository(db)

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		CotacaoDolarHandler(w, r, cotacaoRepo)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func CotacaoDolarHandler(w http.ResponseWriter, r *http.Request, repo *repositorys.CotacaoRepository) {
	ctxHttp := context.Background()
	ctxHttp, cancelHttp := context.WithTimeout(ctxHttp, 200*time.Millisecond)
	defer cancelHttp()

	cotacao, err := rest_clients.BuscarCotacaoUsdParaBrl(ctxHttp)

	if err == nil {
		ctxDB := context.Background()
		ctxDB, cancelDB := context.WithTimeout(ctxDB, 10*time.Millisecond)
		defer cancelDB()

		err = repo.Save(ctxDB, &cotacao.UsdParaBrl)
	}

	var bidResponse response.CotacaoBidResponse

	if cotacao != nil {
		bidResponse.Bid = cotacao.UsdParaBrl.Bid
	}

	response.Responder(w, bidResponse, err)
}
