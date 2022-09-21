package repository

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/vo"
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

type FindAllSubAccountRequest struct {
	WalletID          entity.WalletID
	ParentAccountName string             `form:"parent_account_name,omitempty,default="`
	SubAccountName    string             `form:"sub_account_name,omitempty,default="`
	Side              entity.AccountSide `form:"side,omitempty,default="`
}

type FindAllSubAccountRepo interface {
	FindAllSubAccount(ctx context.Context, page, size int64, req FindAllSubAccountRequest) ([]*entity.SubAccount, int64, error)
}

type FindAllJournalRequest struct {
	WalletID    entity.WalletID
	UserID      entity.UserID `form:"user_id,omitempty,default="`
	Description string        `form:"description,omitempty,default="`
	DateStart   vo.Date       `form:"date_start,omitempty,default="`
	DateEnd     vo.Date       `form:"date_end,omitempty,default="`
}

type FindAllJournalRepo interface {
	FindAllJournal(ctx context.Context, page, size int64, req FindAllJournalRequest) ([]*entity.Journal, int64, error)
}

type FindAllSubAccountBalanceRequest struct {
	WalletID       entity.WalletID
	SubAccountName string  `form:"sub_account_name,omitempty,default="`
	DateStart      vo.Date `form:"date_start,omitempty,default="`
	DateEnd        vo.Date `form:"date_end,omitempty,default="`
}

type FindAllSubAccountBalanceRepo interface {
	FindAllSubAccountBalance(ctx context.Context, page, size int64, req FindAllSubAccountBalanceRequest) ([]*entity.SubAccountBalance, int64, error)
}

type FindLastSubAccountBalancesRepo interface {
	FindLastSubAccountBalances(ctx context.Context, req FindLastSubAccountBalancesRequest) (map[entity.SubAccountCode]entity.Money, error)
}

type FindLastSubAccountBalancesRequest struct {
	WalletID        entity.WalletID
	SubAccountCodes []entity.SubAccountCode
}

type FindAllJournalByIDsRepo interface {
	FindAllJournalByIDs(ctx context.Context, walletId entity.WalletID, journalIDs []entity.JournalID) (map[entity.JournalID]entity.Journal, error)
}

type FindAllSubAccountBalanceByJournalIDRepo interface {
	FindAllSubAccountBalanceByJournalID(ctx context.Context, walletId entity.WalletID, journalIDs []entity.JournalID) (map[entity.JournalID][]entity.SubAccountBalance, error)
}
