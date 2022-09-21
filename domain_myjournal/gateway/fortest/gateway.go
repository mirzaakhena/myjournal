package fortest

import (
	"context"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"myjournal/shared/driver"
	"myjournal/shared/infrastructure/config"
	"myjournal/shared/infrastructure/database"
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

type Something[T any] struct {
	database.AdapterGateway[T]
}

func (s Something[T]) GetAll(param database.GetAllParam, results *[]T) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *gateway) FindAllAccount(ctx context.Context) database.GetAllRepo[entity.Account] {
	r.log.Info(ctx, "called")
	return &Something[entity.Account]{}
}

func (r *gateway) FindAllJournal(ctx context.Context) database.GetAllRepo[entity.Journal] {
	r.log.Info(ctx, "called")
	return &Something[entity.Journal]{}
}

func (r *gateway) FindAllSubAccount(ctx context.Context) database.GetAllRepo[entity.SubAccount] {
	r.log.Info(ctx, "called")
	return &Something[entity.SubAccount]{}
}

func (r *gateway) FindAllSubAccountBalance(ctx context.Context) database.GetAllRepo[entity.SubAccountBalance] {
	r.log.Info(ctx, "called")
	return &Something[entity.SubAccountBalance]{}
}

func (r *gateway) SaveAccounts(ctx context.Context) database.InsertManyRepo[entity.Account] {
	r.log.Info(ctx, "called")
	return &Something[entity.Account]{}
}

func (r *gateway) SaveJournal(ctx context.Context) database.InsertOrUpdateRepo[entity.Journal] {
	r.log.Info(ctx, "called")
	return &Something[entity.Journal]{}
}

func (r *gateway) SaveSubAccountBalances(ctx context.Context) database.InsertManyRepo[entity.SubAccountBalance] {
	r.log.Info(ctx, "called")
	return &Something[entity.SubAccountBalance]{}
}

func (r *gateway) FindLastSubAccountBalances(ctx context.Context, req repository.FindLastSubAccountBalancesRequest) (map[entity.SubAccountCode]entity.Money, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) FindSubAccounts(ctx context.Context) database.GetAllEachItemRepo[entity.SubAccount] {
	r.log.Info(ctx, "called")
	return &Something[entity.SubAccount]{}
}

func (r *gateway) SaveSubAccounts(ctx context.Context) database.InsertManyRepo[entity.SubAccount] {
	r.log.Info(ctx, "called")
	return &Something[entity.SubAccount]{}
}

func (r *gateway) FindAccounts(ctx context.Context) database.GetAllEachItemRepo[entity.Account] {
	r.log.Info(ctx, "called")
	return &Something[entity.Account]{}
}
