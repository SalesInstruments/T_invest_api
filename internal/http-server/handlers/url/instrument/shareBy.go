package instrument

import (
	grpctinvest "T_invest_api/internal/gRPC_TInvest"
	g "T_invest_api/internal/globals"
	"T_invest_api/internal/logger"
	"encoding/json"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ticker := r.URL.Query().Get("ticker")
		classCode := r.URL.Query().Get("classCode")

		if ticker == "" || classCode == "" {
			g.Log.Error("Missing ticker or classCode")
			http.Error(w, "Missing ticker or classCode", http.StatusBadRequest)
			return
		}

		ch := make(chan struct {
			response []byte
			err      error
		})

		go func() {
			defer close(ch)

			conn, err := grpctinvest.New()
			if err != nil {
				ch <- struct {
					response []byte
					err      error
				}{nil, err}
				return
			}
			defer conn.Close()
			defer conn.CancelFunc()

			instrument, err := conn.ShareBy(
				ticker,
				classCode,
			)
			if err != nil {
				ch <- struct {
					response []byte
					err      error
				}{nil, err}
				return
			}

			jsonResponse, err := json.Marshal(instrument)
			if err != nil {
				ch <- struct {
					response []byte
					err      error
				}{nil, err}
				return
			}
			ch <- struct {
				response []byte
				err      error
			}{jsonResponse, nil}
		}()

		result := <-ch

		// Обрабатываем ошибку, если она есть
		if result.err != nil {
			g.Log.Error("Error processing request", logger.Err(result.err))
			http.Error(w, result.err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result.response)
	}
}
