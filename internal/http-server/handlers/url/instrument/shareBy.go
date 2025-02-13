package instrument

import (
	"T_invest_api/internal/config"
	grpctinvest "T_invest_api/internal/gRPC_TInvest"
	"T_invest_api/internal/logger"
	"encoding/json"
	"net/http"
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
			log.Error("Missing ticker or classCode")
			http.Error(w, "Missing ticker or classCode", http.StatusBadRequest)
			return
		}

		conn, err := grpctinvest.New()
		if err != nil {
			log.Error("Failed to create gRPC connection: ", logger.Err(err))
			http.Error(w, "Failed to create gRPC connection: ", http.StatusInternalServerError)
		}
		defer conn.Close()
		defer conn.CancelFunc()

		instrument, err := conn.ShareBy(
			ticker,
			classCode,
		)
		if err != nil {
			log.Error("error receiving the tool", logger.Err(err))
			http.Error(w, "error receiving the tool", http.StatusInternalServerError)
		}

		jsonResponse, err := json.Marshal(instrument)
		if err != nil {
			log.Error("Failed to marshal JSON", logger.Err(err))
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)

	}
}
