package bonds

import (
	"T_invest_api/internal/config"
	grpctinvest "T_invest_api/internal/gRPC_TInvest"
	"T_invest_api/internal/logger"
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

		conn, err := grpctinvest.New()
		if err != nil {
			log.Error("Failed to create gRPC connection: ", logger.Err(err))
			os.Exit(1)
		}
		defer conn.Close()
		defer conn.CancelFunc()

		w.WriteHeader(http.StatusOK)
		// Пишем "Hello, World!" в тело ответа
		w.Write([]byte(conn.Gg()))
	}
}
