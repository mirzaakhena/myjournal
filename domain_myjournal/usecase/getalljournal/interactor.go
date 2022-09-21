package getalljournal

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
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

	objs, count, err := r.outport.FindAllJournal(ctx, req.Page, req.Size, req.FindAllJournalRequest)
	if err != nil {
		return nil, err
	}

	journalIDs := make([]entity.JournalID, 0)
	for _, obj := range objs {
		journalIDs = append(journalIDs, obj.ID)
	}

	if len(journalIDs) == 0 {
		return res, nil
	}

	subAccountBalanceObjs, err := r.outport.FindAllSubAccountBalanceByJournalID(ctx, req.WalletID, journalIDs)
	if err != nil {
		return nil, err
	}

	for _, obj := range objs {

		balances := make([]SubAccountBalance, 0)

		subAccountBalances := subAccountBalanceObjs[obj.ID]

		for _, sab := range subAccountBalances {

			subAccount := sab.SubAccount
			account := subAccount.ParentAccount

			balances = append(balances, SubAccountBalance{
				ID: sab.ID,
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
				Amount:    sab.Amount,
				Balance:   sab.Balance,
				Sequence:  sab.Sequence,
				Direction: sab.Direction,
			})
		}

		res.Items = append(res.Items, Journal{
			ID:          obj.ID,
			Date:        obj.Date,
			WalletID:    obj.WalletID,
			UserID:      obj.UserID,
			Description: obj.Description,
			Balances:    balances,
		})
	}

	res.Count = count

	return res, nil
}
