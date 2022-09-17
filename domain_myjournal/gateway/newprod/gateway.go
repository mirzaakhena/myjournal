package newprod

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"myjournal/shared/driver"
	"myjournal/shared/infrastructure/config"
	"myjournal/shared/infrastructure/logger"
)

type gateway struct {
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
	}
}

func (r *gateway) FindAllAccount(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.Account, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindAllJournal(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.Journal, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindAllSubAccount(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.SubAccount, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindAllSubAccountBalance(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.SubAccountBalance, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) SaveAccounts(ctx context.Context, objs []*entity.Account) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SaveJournal(ctx context.Context, obj *entity.Journal) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SaveSubAccountBalances(ctx context.Context, objs []*entity.SubAccountBalance) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SaveAccountBalances(ctx context.Context, objs []*entity.AccountBalance) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) FindLastSubAccountBalances(ctx context.Context, req repository.FindLastSubAccountBalancesRequest) (map[entity.SubAccountCode]entity.Money, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) FindSubAccounts(ctx context.Context, req repository.FindSubAccountsRequest) (map[entity.SubAccountCode]entity.SubAccount, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveSubAccounts(ctx context.Context, objs []*entity.SubAccount) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) FindAccounts(ctx context.Context, req repository.FindAccountsRequest) (map[entity.AccountCode]entity.Account, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}
