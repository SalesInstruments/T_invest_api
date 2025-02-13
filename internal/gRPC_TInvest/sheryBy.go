package grpctinvest

import (
	pb "T_invest_api/internal/gRPC_TInvest/proto"
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
	if err != nil {
		return nil, err
	}

	return r, nil
}
