package repositorys

import (
	"context"
	"database/sql"
	"server/domain/Entity"
	"server/domain/errors"
)

type CotacaoRepository struct {
	db *sql.DB
}

func NewCotacaoRepository(db *sql.DB) *CotacaoRepository {
	return &CotacaoRepository{db: db}
}

func (r *CotacaoRepository) Save(ctx context.Context, cotacao *Entity.Cotacao) *errors.CustomError {
	query := `INSERT INTO cotacoes (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, cotacao.Code, cotacao.Codein, cotacao.Name, cotacao.High, cotacao.Low, cotacao.VarBid, cotacao.PctChange, cotacao.Bid, cotacao.Ask, cotacao.Timestamp, cotacao.CreateDate)

	if err != nil {
		errCreate := errors.NewCotacaoRepositoryCreateError(err)
		return &errCreate
	}

	return nil
}
