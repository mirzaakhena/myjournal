package getallaccountbalance

import (
	"context"
)

//go:generate mockery --name Outport -output mocks/

type getAllAccountBalanceInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllAccountBalanceInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllAccountBalanceInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	//!

	return res, nil
}
