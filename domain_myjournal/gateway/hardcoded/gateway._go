package hardcoded

import (
	"context"
	"encoding/json"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"myjournal/shared/driver"
	"myjournal/shared/infrastructure/config"
	"myjournal/shared/infrastructure/logger"
	"myjournal/shared/util"
	"os"
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

func (r *gateway) SaveAccounts(ctx context.Context, objs []*entity.Account) error {
	r.log.Info(ctx, "SaveAccounts %v", util.MustJSON(objs))

	return nil
}

func (r *gateway) SaveJournal(ctx context.Context, obj *entity.Journal) error {
	r.log.Info(ctx, "SaveJournal %s", util.MustJSON(obj))

	return nil
}

func (r *gateway) SaveSubAccountBalances(ctx context.Context, objs []*entity.SubAccountBalance) error {
	r.log.Info(ctx, "SaveSubAccountBalances %s", util.MustJSON(objs))

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

	dataBytes, err := os.ReadFile("domain_myjournal/gateway/prod/subaccounts.json")
	if err != nil {
		return nil, err
	}

	dataObjs := make([]entity.SubAccount, 0)
	err = json.Unmarshal(dataBytes, &dataObjs)
	if err != nil {
		return nil, err
	}

	results := map[entity.SubAccountCode]entity.SubAccount{}

	for _, account := range dataObjs {
		results[account.Code] = account
	}

	return results, nil
}

func (r *gateway) GetRandomString(ctx context.Context) string {
	r.log.Info(ctx, "called")
	return util.GenerateID(5)
}

func (r *gateway) SaveSubAccounts(ctx context.Context, objs []*entity.SubAccount) error {
	r.log.Info(ctx, "SaveSubAccounts %v", util.MustJSON(objs))

	return nil
}

func (r *gateway) FindAccounts(ctx context.Context, req repository.FindAccountsRequest) (map[entity.AccountCode]entity.Account, error) {
	r.log.Info(ctx, "called")

	dataBytes, err := os.ReadFile("domain_myjournal/gateway/prod/accounts.json")
	if err != nil {
		return nil, err
	}

	dataObjs := make([]entity.Account, 0)
	err = json.Unmarshal(dataBytes, &dataObjs)
	if err != nil {
		return nil, err
	}

	results := map[entity.AccountCode]entity.Account{}

	for _, account := range dataObjs {
		results[account.Code] = account
	}

	return results, nil
}
