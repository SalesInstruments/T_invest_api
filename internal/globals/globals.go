package globals

import (
	"T_invest_api/internal/config"
	"T_invest_api/internal/logger"
)

var (
	Cfg             = config.MustLoad()
	CfgPostgresDB   = Cfg.PostgresDB
	CfgRedisDB      = Cfg.RedisDB
	CfgGRPS_TInvest = Cfg.GRPC_TInvest_server
	Log             = logger.SetupLogger(Cfg.Env)
)
