package grpctinvest

import (
	pb "T_invest_api/internal/gRPC_TInvest/proto"
	"T_invest_api/internal/logger"
	"fmt"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc/status"
)

func (grpcc *GRPCconn) ShareBy(ticker, classCode string) (*pb.ShareResponse, error) {
	c := pb.NewInstrumentsServiceClient(grpcc.ClientConn)

	r, err := c.ShareBy(
		grpcc.Context,
		&pb.InstrumentRequest{
			IdType:    pb.InstrumentIdType_INSTRUMENT_ID_TYPE_TICKER,
			ClassCode: &classCode,
			Id:        "SBER",
		},
	)

	fmt.Println(r)

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Error("gRPC error", slog.String("message", st.Message()), slog.Any("details", st.Details()))
		} else {
			log.Error("could not get instrument", logger.Err(err))
		}
		return nil, err
	}

	return r, nil
}
