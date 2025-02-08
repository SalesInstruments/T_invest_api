package storage

import (
	"T_invest_api/internal/config"
	"T_invest_api/internal/logger"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	cfg   = config.MustLoad()
	cfgDB = cfg.PostgresDB
	log   = logger.SetupLogger(cfg.Env)
)

type PostgresDB struct {
	*sqlx.DB
}

func New() (*PostgresDB, error) {

	log.Debug(
		"PostgresDB params:",
		slog.String("host", cfgDB.Host),
		slog.String("port", strconv.Itoa(cfgDB.Port)),
		slog.String("username", cfgDB.Username),
		slog.String("password", cfgDB.Password),
		slog.String("dbname", cfgDB.DBname),
		slog.String("sslmode", cfgDB.SSLmode),
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfgDB.Host, cfgDB.Port, cfgDB.Username, cfgDB.Password, cfgDB.DBname, cfgDB.SSLmode,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Info("successfully connected to the database")

	return &PostgresDB{db}, nil
}

func (db *PostgresDB) Close() error {
	log.Info("closing database connection")
	return db.DB.Close()
}
