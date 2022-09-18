package getallsubaccount

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/shared/infrastructure/database"
	"myjournal/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type getAllSubAccountInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllSubAccountInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllSubAccountInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	//subAccountObjs, count, err := r.outport.FindAllSubAccount(ctx, req.Page, req.Size, req.WalletId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//res.Count = count
	//res.Items = util.ToSliceAny(subAccountObjs)

	p := database.NewDefaultParam().
		Page(req.Page).
		Size(req.Size).
		Filter("parent_account.wallet_id", req.WalletId).
		Sort("code", 1)

	objs := make([]entity.SubAccount, 0)
	count, err := r.outport.FindAllSubAccount(ctx).GetAll(p, &objs)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(objs)

	return res, nil
}
