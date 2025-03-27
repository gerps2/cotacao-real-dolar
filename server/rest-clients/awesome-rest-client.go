package rest_clients

import (
	errorsGo "errors"

	"context"
	"encoding/json"
	"io"
	"net/http"
	"server/domain/errors"
	"server/domain/response"
)

func BuscarCotacaoUsdParaBrl(ctx context.Context) (*response.CotacaoResponse, *errors.CustomError) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		customErr := errors.NewRequestError(err)
		return nil, &customErr
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		if errorsGo.Is(err, context.DeadlineExceeded) {
			customErr := errors.NewTimeoutError(err)
			return nil, &customErr
		}

		customErr := errors.NewRequestError(err)
		return nil, &customErr
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		customErr := errors.NewStatusCodeError(err)
		return nil, &customErr
	}

	var cotacaoResponse response.CotacaoResponse
	err = json.Unmarshal(body, &cotacaoResponse)
	if err != nil {
		customErr := errors.NewUnmarshalError(err)
		return nil, &customErr
	}

	return &cotacaoResponse, nil
}
