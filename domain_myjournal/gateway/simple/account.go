package simple

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"myjournal/shared/infrastructure/database"
	"myjournal/shared/infrastructure/logger"
)

type accountRepoImpl struct {
	log  logger.Logger
	repo database.Repository[entity.Account]
}

func NewAccountRepoImpl(log logger.Logger, db *mongo.Database) *accountRepoImpl {
	return &accountRepoImpl{
		log:  log,
		repo: database.NewMongoGateway[entity.Account](db),
	}
}

func (r accountRepoImpl) SaveAccounts(ctx context.Context, objs []*entity.Account) error {
	r.log.Info(ctx, "called")

	err := r.repo.InsertMany(objs...)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return err
	}

	return nil
}

func (r accountRepoImpl) FindAllAccount(ctx context.Context, page, size int64, walletId entity.WalletID) ([]*entity.Account, int64, error) {
	r.log.Info(ctx, "called")

	p := database.NewDefaultParam().
		Page(page).
		Size(size).
		Filter("wallet_id", walletId).
		Sort("code", 1)

	objs := make([]*entity.Account, 0)
	count, err := r.repo.GetAll(p, &objs)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return objs, count, nil
}

func (r accountRepoImpl) FindAccounts(ctx context.Context, req repository.FindAccountsRequest) (map[entity.AccountCode]entity.Account, error) {
	r.log.Info(ctx, "called")

	var bsonObjectIDs bson.A
	for _, val := range req.AccountIds {
		bsonObjectIDs = append(bsonObjectIDs, val)
	}

	p := database.NewDefaultParam().
		Page(1).
		Size(100).
		Filter("wallet_id", req.WalletID).
		Filter("_id", bson.M{"$in": bsonObjectIDs}).
		Sort("code", 1)

	results := map[entity.AccountCode]entity.Account{}
	_, err := r.repo.GetAllEachItem(p, func(obj entity.Account) {
		results[obj.Code] = obj
	})
	if err != nil {
		r.log.Error(ctx, err.Error())
		return nil, err
	}

	return results, nil
}
