package postgresdb

import (
	g "T_invest_api/internal/globals"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	*sqlx.DB
}

func New() (*PostgresDB, error) {

	g.Log.Debug(
		"PostgresDB params:",
		slog.String("host", g.CfgPostgresDB.Host),
		slog.Int("port", g.CfgPostgresDB.Port),
		slog.String("username", g.CfgPostgresDB.Username),
		slog.String("password", g.CfgPostgresDB.Password),
		slog.String("dbname", g.CfgPostgresDB.DBname),
		slog.String("sslmode", g.CfgPostgresDB.SSLmode),
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		g.CfgPostgresDB.Host,
		g.CfgPostgresDB.Port,
		g.CfgPostgresDB.Username,
		g.CfgPostgresDB.Password,
		g.CfgPostgresDB.DBname,
		g.CfgPostgresDB.SSLmode,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	g.Log.Info("successfully connected to the postgres database")

	return &PostgresDB{db}, nil
}

func (db *PostgresDB) Close() error {
	g.Log.Info("closing database connection")
	return db.DB.Close()
}
