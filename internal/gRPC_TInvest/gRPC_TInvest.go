package grpctinvest

import (
	g "T_invest_api/internal/globals"
	"T_invest_api/internal/logger"
	"context"
	"crypto/tls"
	"time"

	"golang.org/x/exp/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type GRPCconn struct {
	*grpc.ClientConn
	context.Context
	context.CancelFunc
}

func New() (*GRPCconn, error) {
	// token := &tokenAuth{token: cfgGRPS_TInvest.Token}

	ctx := context.Background()

	g.Log.Debug(
		"gRPC connect params",
		slog.String("address", g.CfgGRPS_TInvest.SAddress),
		slog.String("token", g.CfgGRPS_TInvest.Token),
	)

	g.Log.Info("try connect gRPC")

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	conn, err := grpc.NewClient(
		g.CfgGRPS_TInvest.SAddress,
		grpc.WithTransportCredentials(creds),
	)

	if err != nil {
		g.Log.Error("did not connect: ", logger.Err(err))
		return nil, err
	}
	g.Log.Info(" connect gRPC")

	md := metadata.Pairs("authorization", "Bearer "+g.CfgGRPS_TInvest.Token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	return &GRPCconn{
		conn,
		ctx,
		cancel,
	}, nil
}

// func (grpcc *GRPCconn) Gg() string {
// 	c := pb.NewInstrumentsServiceClient(grpcc.ClientConn)

// 	// classCode := "TQBR"
// 	instrumentStatus := pb.InstrumentStatus_INSTRUMENT_STATUS_ALL

// 	r, err := c.Shares(
// 		grpcc.Context,
// 		&pb.InstrumentsRequest{
// 			InstrumentStatus: &instrumentStatus,
// 			// IdType:    pb.InstrumentIdType_INSTRUMENT_ID_TYPE_TICKER,
// 			// ClassCode: &classCode,
// 			// Id:        "TCSG",
// 		},
// 	)
// 	if err != nil {
// 		st, ok := status.FromError(err)
// 		if ok {
// 			log.Error("gRPC error", slog.String("message", st.Message()), slog.Any("details", st.Details()))
// 		} else {
// 			log.Error("could not get instrument", logger.Err(err))
// 		}
// 		return ""
// 	}

// 	tiker := "GAZP"

// 	for _, i := range r.Instruments {
// 		if i.Ticker == tiker {
// 			return fmt.Sprintf("имя акции: %s\nтикер акции: %s\nisin фкции: %s\nclass_code фкции: %s",
// 				i.Name,
// 				i.Ticker,
// 				i.Isin,
// 				i.ClassCode)
// 		}
// 	}

// 	fmt.Println(len(r.Instruments))
// 	return r.Instruments[0].String()
// }
