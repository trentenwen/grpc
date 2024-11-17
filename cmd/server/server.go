package server

import (
	"context"

	pb "github.com/trentenwen/grpc/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	pb.UnimplementedCurrencyServer
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{UnimplementedCurrencyServer: pb.UnimplementedCurrencyServer{}, log: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *pb.RateRequest) (*pb.RateResponse, error) {
	c.log.Info("Handle GetRate", "Base", rr.GetBase(), "Dest", rr.GetDestination())

	return &pb.RateResponse{Rate: 0.5}, nil
}
