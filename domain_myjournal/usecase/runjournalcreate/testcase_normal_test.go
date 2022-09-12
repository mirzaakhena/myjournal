package runjournalcreate

import (
	"context"
	"github.com/stretchr/testify/assert"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"myjournal/shared/infrastructure/util"
	"testing"
	"time"
)

type mockOutportNormal struct {
	t   *testing.T
	Now time.Time
}

// TestCaseNormal is for the case where ...
// explain the purpose of this test here with human readable naration...
func TestCaseNormal(t *testing.T) {

	now := time.Now()

	mockDatasource := &mockOutportNormal{
		t:   t,
		Now: now,
		// shared data to outport
	}

	res, err := NewUsecase(mockDatasource).Execute(context.Background(), InportRequest{
		Date:     time.Now(),
		WalletId: "WL01",
		UserId:   "US01",
		Journals: []Journal{
			{
				Description: "Desc01",
				Transactions: []Transaction{
					{Sign: "+", SubAccountCode: "HR001", Amount: 5000},
					{Sign: "+", SubAccountCode: "UT001", Amount: 5000},
				},
			},
		},
	})

	assert.Nil(t, err)

	t.Logf("%v", res)

}

func (r *mockOutportNormal) SaveJournal(ctx context.Context, objs []*entity.Journal) error {
	r.t.Logf("Save Journal %v", util.MustJSON(objs))

	return nil
}

func (r *mockOutportNormal) SaveSubAccountBalances(ctx context.Context, objs []*entity.SubAccountBalance) error {
	r.t.Logf("Save Transactions %v", util.MustJSON(objs))

	assert.Equal(r.t, entity.Money(6200+5000), objs[0].Balance)
	assert.Equal(r.t, entity.Money(12500+5000), objs[1].Balance)
	return nil
}

func (r *mockOutportNormal) SaveAccountBalances(ctx context.Context, objs []*entity.AccountBalance) error {
	r.t.Logf("Save AccountBalances %v", util.MustJSON(objs))

	return nil
}

func (r *mockOutportNormal) FindLastSubAccountBalances(ctx context.Context, req repository.FindLastSubAccountBalancesRepoRequest) (map[entity.SubAccountCode]*entity.SubAccountBalance, error) {
	r.t.Logf("Transactions %v", util.MustJSON(req))

	return map[entity.SubAccountCode]*entity.SubAccountBalance{
		"HR001": {
			Id:        "",
			JournalId: "",
			SubAccount: entity.SubAccount{
				Id:   "",
				Code: "HR001",
				Name: "Harta",
				ParentAccount: entity.Account{
					Id:       "",
					WalletId: "",
					Code:     "HARTA",
					Name:     "Harta",
					Level:    0,
					Side:     entity.AccountSideActiva,
				},
			},
			Date:      r.Now,
			Amount:    0,
			Balance:   6200,
			Sequence:  0,
			Direction: "",
		},
		"UT001": {
			Id:        "",
			JournalId: "",
			SubAccount: entity.SubAccount{
				Id:   "",
				Code: "UT001",
				Name: "Utang",
				ParentAccount: entity.Account{
					Id:       "",
					WalletId: "",
					Code:     "UTANG",
					Name:     "Utang",
					Level:    0,
					Side:     entity.AccountSidePassiva,
				},
			},
			Date:      r.Now,
			Amount:    0,
			Balance:   12500,
			Sequence:  0,
			Direction: "",
		},
	}, nil
}

func (r *mockOutportNormal) FindSubAccounts(ctx context.Context, objs []entity.SubAccountCode) (map[entity.SubAccountCode]entity.SubAccount, error) {
	r.t.Logf("SubAccounts %v", util.MustJSON(objs))

	return map[entity.SubAccountCode]entity.SubAccount{
		"HR001": {
			Id:   "",
			Code: "HR001",
			Name: "Harta",
			ParentAccount: entity.Account{
				Id:       "",
				WalletId: "",
				Code:     "HARTA",
				Name:     "Harta",
				Level:    0,
				Side:     entity.AccountSideActiva,
			},
		},
		"UT001": {
			Id:   "",
			Code: "UT001",
			Name: "Utang",
			ParentAccount: entity.Account{
				Id:       "",
				WalletId: "",
				Code:     "UTANG",
				Name:     "Utang",
				Level:    0,
				Side:     entity.AccountSidePassiva,
			},
		},
	}, nil
}
