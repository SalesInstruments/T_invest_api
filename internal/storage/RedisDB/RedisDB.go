package redisdb

import (
	g "T_invest_api/internal/globals"
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	*redis.Client
}

func New() (*RedisDB, error) {
	g.Log.Debug(
		"RedisDB params:",
		slog.String("address", g.CfgRedisDB.Address),
		slog.String("password", g.CfgRedisDB.Password),
		slog.String("username", g.CfgRedisDB.Username),
		slog.Int("DB", g.CfgRedisDB.DB),
		slog.Int("max_retries", g.CfgRedisDB.MaxRetries),
	)

	db := redis.NewClient(&redis.Options{
		Addr:         g.CfgRedisDB.Address,
		Password:     g.CfgRedisDB.Password,
		DB:           g.CfgRedisDB.DB,
		Username:     g.CfgRedisDB.Username,
		MaxRetries:   g.CfgRedisDB.MaxRetries,
		DialTimeout:  g.CfgRedisDB.DialTimeout,
		ReadTimeout:  g.CfgRedisDB.Timeout,
		WriteTimeout: g.CfgRedisDB.Timeout,
	})

	if err := db.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	g.Log.Info("successfully connected to the redis database")

	return &RedisDB{db}, nil
}
