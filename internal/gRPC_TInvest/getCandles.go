package grpctinvest

import (
	pb "T_invest_api/internal/gRPC_TInvest/proto"
	"T_invest_api/internal/logger"
	"time"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (grpcc *GRPCconn) GetCandles(figi string) (*pb.GetCandlesResponse, error) {
	c := pb.NewMarketDataServiceClient(grpcc.ClientConn)

	instrumentID := "" // Оставляем пустым, если используем figi

	// Указываем временные рамки для запроса
	from := time.Now().Add(-24 * time.Hour) // Начало периода (24 часа назад)
	to := time.Now()                        // Конец периода (текущее время)

	// Преобразуем time.Time в timestamppb.Timestamp
	fromTimestamp := timestamppb.New(from)
	toTimestamp := timestamppb.New(to)

	// Указываем интервал свечей (например, 2 часа)
	interval := pb.CandleInterval_CANDLE_INTERVAL_1_MIN

	// Указываем тип источника свечи (опционально)
	candleSourceType := pb.GetCandlesRequest_CANDLE_SOURCE_EXCHANGE

	r, err := c.GetCandles(
		grpcc.Context,
		&pb.GetCandlesRequest{
			Figi:             &figi,
			From:             fromTimestamp,     // Начало периода
			To:               toTimestamp,       // Конец периода
			Interval:         interval,          // Интервал свечей
			InstrumentId:     &instrumentID,     // Указываем instrument_id (опционально, если используется figi)
			CandleSourceType: &candleSourceType, // Тип источника свечи (опционально)
		},
	)

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Error("gRPC error", slog.String("message", st.Message()), slog.Any("details", st.Details()))
		} else {
			log.Error("could not get candles", logger.Err(err))
		}
		return nil, err
	}

	return r, nil
}
