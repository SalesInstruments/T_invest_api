package grpctinvest

import (
	"T_invest_api/internal/config"
	"T_invest_api/internal/logger"
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"time"

	pb "T_invest_api/internal/gRPC_TInvest/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	cfg             = config.MustLoad()
	cfgGRPS_TInvest = cfg.GRPC_TInvest_server
	log             = logger.SetupLogger(cfg.Env)
)

// type tokenAuth struct {
// 	token string
// }

// func (t *tokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
// 	return map[string]string{
// 		"authorization": "Bearer " + t.token,
// 	}, nil
// }

// func (t *tokenAuth) RequireTransportSecurity() bool {
// 	return true // Использовать TLS для безопасной передачи токена
// }

type GRPCconn struct {
	*grpc.ClientConn
	context.Context
	context.CancelFunc
}

func New() (*GRPCconn, error) {
	// token := &tokenAuth{token: cfgGRPS_TInvest.Token}

	ctx := context.Background()

	log.Debug(
		"gRPC connect params:",
		slog.String("address", cfgGRPS_TInvest.SAddress),
		slog.String("token", cfgGRPS_TInvest.Token),
	)

	log.Info("try connect gRPC")

	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	conn, err := grpc.NewClient(
		cfgGRPS_TInvest.SAddress,
		grpc.WithTransportCredentials(creds),
	)

	if err != nil {
		log.Error("did not connect: ", logger.Err(err))
		return nil, err
	}
	log.Info(" connect gRPC")

	md := metadata.Pairs("authorization", "Bearer "+cfgGRPS_TInvest.Token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	return &GRPCconn{
		conn,
		ctx,
		cancel,
	}, nil
}

func (grpcc *GRPCconn) Gg() string {
	c := pb.NewInstrumentsServiceClient(grpcc.ClientConn)
	r, err := c.BondBy(
		grpcc.Context,
		&pb.InstrumentRequest{
			IdType: pb.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
			Id:     "BBG00QXGFHS6",
		},
	)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Error("gRPC error", slog.String("message", st.Message()), slog.Any("details", st.Details()))
		} else {
			log.Error("could not get instrument", logger.Err(err))
		}
		return ""
	}

	fmt.Println(r.Instrument.Name)
	return r.Instrument.Name
}
