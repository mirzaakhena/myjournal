package repository

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/shared/infrastructure/database"
)

//type SaveJournalRepo interface {
//	SaveJournal(ctx context.Context, obj *entity.Journal) error
//}

type SaveJournalRepo2 interface {
	SaveJournal(ctx context.Context) database.InsertOrUpdateRepo[entity.Journal]
}

//type SaveSubAccountBalancesRepo interface {
//	SaveSubAccountBalances(ctx context.Context, objs []*entity.SubAccountBalance) error
//}

type SaveSubAccountBalancesRepo2 interface {
	SaveSubAccountBalances(ctx context.Context) database.InsertManyRepo[entity.SubAccountBalance]
}

//type SaveAccountBalancesRepo interface {
//	SaveAccountBalances(ctx context.Context, objs []*entity.AccountBalance) error
//}

type SaveAccountBalancesRepo2 interface {
	SaveAccountBalances(ctx context.Context) database.InsertManyRepo[entity.AccountBalance]
}

//type SaveAccountsRepo interface {
//	SaveAccounts(ctx context.Context, objs []*entity.Account) error
//}

type SaveAccountsRepo2 interface {
	SaveAccounts(ctx context.Context) database.InsertManyRepo[entity.Account]
}

//type SaveSubAccountsRepo interface {
//	SaveSubAccounts(ctx context.Context, objs []*entity.SubAccount) error
//}

type SaveSubAccountsRepo2 interface {
	SaveSubAccounts(ctx context.Context) database.InsertManyRepo[entity.SubAccount]
}

//type FindAccountsRepo interface {
//	FindAccounts(ctx context.Context, req FindAccountsRequest) (map[entity.AccountCode]entity.Account, error)
//}

type FindAccountsRepo2 interface {
	FindAccounts(ctx context.Context) database.GetAllEachItemRepo[entity.Account]
}

//type FindAccountsRequest struct {
//	WalletID   entity.WalletID
//	AccountIds []entity.AccountID
//}

type FindLastSubAccountBalancesRepo interface {
	FindLastSubAccountBalances(ctx context.Context, req FindLastSubAccountBalancesRequest) (map[entity.SubAccountCode]entity.Money, error)
}

type FindLastSubAccountBalancesRequest struct {
	WalletID        entity.WalletID
	SubAccountCodes []entity.SubAccountCode
}

//type FindSubAccountsRepo interface {
//	FindSubAccounts(ctx context.Context, req FindSubAccountsRequest) (map[entity.SubAccountCode]entity.SubAccount, error)
//}

type FindSubAccountsRepo2 interface {
	FindSubAccounts(ctx context.Context) database.GetAllEachItemRepo[entity.SubAccount]
}

//type FindSubAccountsRequest struct {
//	WalletID        entity.WalletID
//	SubAccountCodes []entity.SubAccountCode
//}

//type FindAllAccountRepo interface {
//	FindAllAccount(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.Account, int64, error)
//}

type FindAllAccountRepo2 interface {
	FindAllAccount(ctx context.Context) database.GetAllRepo[entity.Account]
}

//type FindAllSubAccountRepo interface {
//	FindAllSubAccount(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.SubAccount, int64, error)
//}

type FindAllSubAccountRepo2 interface {
	FindAllSubAccount(ctx context.Context) database.GetAllRepo[entity.SubAccount]
}

//type FindAllJournalRepo interface {
//	FindAllJournal(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.Journal, int64, error)
//}

type FindAllJournalRepo2 interface {
	FindAllJournal(ctx context.Context) database.GetAllRepo[entity.Journal]
}

//type FindAllSubAccountBalanceRepo interface {
//	FindAllSubAccountBalance(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.SubAccountBalance, int64, error)
//}

type FindAllSubAccountBalanceRepo2 interface {
	FindAllSubAccountBalance(ctx context.Context) database.GetAllRepo[entity.SubAccountBalance]
}
