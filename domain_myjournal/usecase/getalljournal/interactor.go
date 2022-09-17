package getalljournal

import (
	"context"
	"myjournal/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type getAllJournalInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllJournalInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllJournalInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	journalObjs, count, err := r.outport.FindAllJournal(ctx, req.Page, req.Size, req.WalletId)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(journalObjs)

	return res, nil
}
