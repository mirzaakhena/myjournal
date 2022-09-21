package getallaccount

import (
	"context"
	"myjournal/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type getAllAccountInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllAccountInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllAccountInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	objs, count, err := r.outport.FindAllAccount(ctx, req.Page, req.Size, req.WalletId)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(objs)

	return res, nil
}
