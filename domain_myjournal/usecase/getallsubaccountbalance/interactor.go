package getallsubaccountbalance

import (
	"context"
	"myjournal/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type getAllSubaccountBalanceInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllSubaccountBalanceInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllSubaccountBalanceInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	subAccountBalanceObjs, count, err := r.outport.FindAllSubAccountBalance(ctx, req.Page, req.Size, req.WalletId)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(subAccountBalanceObjs)

	return res, nil
}
