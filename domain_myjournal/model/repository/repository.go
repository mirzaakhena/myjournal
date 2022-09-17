package repository

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
)

type SaveJournalRepo interface {
	SaveJournal(ctx context.Context, obj *entity.Journal) error
}

type SaveSubAccountBalancesRepo interface {
	SaveSubAccountBalances(ctx context.Context, objs []*entity.SubAccountBalance) error
}

type SaveAccountBalancesRepo interface {
	SaveAccountBalances(ctx context.Context, objs []*entity.AccountBalance) error
}

type SaveAccountsRepo interface {
	SaveAccounts(ctx context.Context, objs []*entity.Account) error
}

type SaveSubAccountsRepo interface {
	SaveSubAccounts(ctx context.Context, objs []*entity.SubAccount) error
}

type FindAccountsRepo interface {
	FindAccounts(ctx context.Context, req FindAccountsRequest) (map[entity.AccountCode]entity.Account, error)
}

type FindAccountsRequest struct {
	WalletID   entity.WalletID
	AccountIds []entity.AccountID
}

type FindLastSubAccountBalancesRepo interface {
	FindLastSubAccountBalances(ctx context.Context, req FindLastSubAccountBalancesRequest) (map[entity.SubAccountCode]entity.Money, error)
}

type FindLastSubAccountBalancesRequest struct {
	WalletID        entity.WalletID
	SubAccountCodes []entity.SubAccountCode
}

type FindSubAccountsRepo interface {
	FindSubAccounts(ctx context.Context, req FindSubAccountsRequest) (map[entity.SubAccountCode]entity.SubAccount, error)
}

type FindSubAccountsRequest struct {
	WalletID        entity.WalletID
	SubAccountCodes []entity.SubAccountCode
}

type FindAllAccountRepo interface {
	FindAllAccount(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.Account, int64, error)
}

type FindAllSubAccountRepo interface {
	FindAllSubAccount(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.SubAccount, int64, error)
}

type FindAllJournalRepo interface {
	FindAllJournal(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.Journal, int64, error)
}

type FindAllSubAccountBalanceRepo interface {
	FindAllSubAccountBalance(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.SubAccountBalance, int64, error)
}
