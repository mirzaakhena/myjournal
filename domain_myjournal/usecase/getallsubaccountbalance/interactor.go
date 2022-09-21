package getallsubaccountbalance

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
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

	objs, count, err := r.outport.FindAllSubAccountBalance(ctx, req.Page, req.Size, req.FindAllSubAccountBalanceRequest)
	if err != nil {
		return nil, err
	}

	journalIDs := make([]entity.JournalID, 0)
	for _, obj := range objs {
		journalIDs = append(journalIDs, obj.JournalID)
	}

	journalObjs, err := r.outport.FindAllJournalByIDs(ctx, req.WalletID, journalIDs)
	if err != nil {
		return nil, err
	}

	res.Count = count

	for _, obj := range objs {

		subAccount := obj.SubAccount
		account := subAccount.ParentAccount
		journal := journalObjs[obj.JournalID]

		res.Items = append(res.Items, TheItem{
			ID:       obj.ID,
			WalletID: journal.WalletID,
			Journal: Journal{
				ID:          journal.ID,
				UserID:      journal.UserID,
				Description: journal.Description,
			},
			SubAccount: SubAccount{
				ID:   subAccount.ID,
				Code: subAccount.Code,
				Name: subAccount.Name,
				ParentAccount: Account{
					ID:    account.ID,
					Code:  account.Code,
					Name:  account.Name,
					Level: account.Level,
					Side:  account.Side,
				},
			},
			Date:      obj.Date,
			Amount:    obj.Amount,
			Balance:   obj.Balance,
			Sequence:  obj.Sequence,
			Direction: obj.Direction,
		})
	}

	return res, nil
}
