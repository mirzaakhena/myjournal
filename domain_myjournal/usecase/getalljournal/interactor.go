package getalljournal

import (
	"context"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
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

	p := message.NewPrinter(language.Indonesian)

	for _, obj := range objs {

		balances := make([]SubAccountBalance, 0)

		subAccountBalances := subAccountBalanceObjs[obj.ID]

		for _, sab := range subAccountBalances {

			amountDebit := ""
			amountCredit := ""

			if sab.Direction == entity.BalanceDirectionDebit {

				amountDebit = p.Sprintf("%.2f", math.Abs(float64(sab.Amount)))

				//amountDebit = fmt.Sprintf("%.000f", math.Abs(float64(sab.Amount)))
			} else if sab.Direction == entity.BalanceDirectionCredit {

				amountCredit = p.Sprintf("%.2f", math.Abs(float64(sab.Amount)))
				//amountCredit = fmt.Sprintf("%.000f", math.Abs(float64(sab.Amount)))
			}

			balances = append(balances, SubAccountBalance{
				ID:             sab.ID,
				AmountDebit:    amountDebit,
				AmountCredit:   amountCredit,
				Sequence:       sab.Sequence,
				SubAccountCode: sab.SubAccount.Code,
				SubAccountName: sab.SubAccount.Name,
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
