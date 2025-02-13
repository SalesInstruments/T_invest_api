package instrument

import (
	"T_invest_api/internal/config"
	grpctinvest "T_invest_api/internal/gRPC_TInvest"
	"T_invest_api/internal/logger"
	"encoding/json"
	"net/http"
	"os"
	// "github.com/tinkoff/invest-api-go-sdk/investgo"
)

var (
	cfg = config.MustLoad()
	log = logger.SetupLogger(cfg.Env)
)

// type Request struct {
// 	id_type
// 	class_code string
// 	id         string
// }

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ticker := r.URL.Query().Get("ticker")
		classCode := r.URL.Query().Get("classCode")

		if ticker == "" || classCode == "" {
			http.Error(w, "Missing ticker or classCode", http.StatusBadRequest)
			return
		}

		conn, err := grpctinvest.New()
		if err != nil {
			log.Error("Failed to create gRPC connection: ", logger.Err(err))
			os.Exit(1)
		}
		defer conn.Close()
		defer conn.CancelFunc()

		instrument, err := conn.ShareBy(
			ticker,
			classCode,
		)
		if err != nil {
			log.Error("error receiving the tool", logger.Err(err))
		}

		jsonResponse, err := json.Marshal(instrument)
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			// return nil, err
		}

		// return jsonResponse, nil

		w.Header().Set("Content-Type", "application/json")

		// Записываем JSON в тело ответа
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)

	}
}
