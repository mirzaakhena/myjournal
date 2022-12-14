package runjournalcreate

import (
	"context"
	"github.com/stretchr/testify/assert"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"myjournal/shared/util"
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
		Date:        time.Now(),
		WalletId:    "WL01",
		UserId:      "US01",
		Description: "Desc01",
		Transactions: []Transaction{
			{Sign: "+", SubAccountCode: "HR001", Amount: 5000},
			{Sign: "+", SubAccountCode: "UT001", Amount: 5000},
		},
	})

	assert.Nil(t, err)

	t.Logf("%v", res)

}

func (r *mockOutportNormal) SaveJournal(ctx context.Context, obj *entity.Journal) error {
	r.t.Logf("Save Journal %v", util.MustJSON(obj))

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

func (r *mockOutportNormal) FindLastSubAccountBalances(ctx context.Context, req repository.FindLastSubAccountBalancesRequest) (map[entity.SubAccountCode]entity.Money, error) {
	r.t.Logf("Transactions %v", util.MustJSON(req))

	return map[entity.SubAccountCode]entity.Money{
		"HR001": 20000,
		"UT001": 32000,
	}, nil
}

func (r *mockOutportNormal) FindSubAccounts(ctx context.Context, req repository.FindSubAccountsRequest) (map[entity.SubAccountCode]entity.SubAccount, error) {
	r.t.Logf("SubAccounts %v", util.MustJSON(req))

	return map[entity.SubAccountCode]entity.SubAccount{
		"HR001": {
			ID:   "",
			Code: "HR001",
			Name: "Harta",
			ParentAccount: entity.Account{
				ID:       "",
				WalletId: "",
				Code:     "HARTA",
				Name:     "Harta",
				Level:    0,
				Side:     entity.AccountSideActiva,
			},
		},
		"UT001": {
			ID:   "",
			Code: "UT001",
			Name: "Utang",
			ParentAccount: entity.Account{
				ID:       "",
				WalletId: "",
				Code:     "UTANG",
				Name:     "Utang",
				Level:    0,
				Side:     entity.AccountSidePassiva,
			},
		},
	}, nil
}
