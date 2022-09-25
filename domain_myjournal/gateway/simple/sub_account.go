package simple

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"myjournal/shared/infrastructure/database"
	"myjournal/shared/infrastructure/logger"
	"strings"
)

type subAccountRepoImpl struct {
	log  logger.Logger
	repo database.Repository[entity.SubAccount]
}

func NewSubAccountRepoImpl(log logger.Logger, db *mongo.Database) *subAccountRepoImpl {
	return &subAccountRepoImpl{
		log:  log,
		repo: database.NewMongoGateway[entity.SubAccount](db),
	}
}

func (r subAccountRepoImpl) FindAllSubAccount(ctx context.Context, page, size int64, req repository.FindAllSubAccountRequest) ([]*entity.SubAccount, int64, error) {

	p := database.NewDefaultParam().
		Page(page).
		Size(size).
		Filter("parent_account.wallet_id", req.WalletID).
		Sort("code", 1)

	if len(req.ParentAccountName) > 0 {
		p = p.Filter("parent_account.name", primitive.Regex{Pattern: req.ParentAccountName, Options: "i"})
	}

	if len(req.SubAccountName) > 0 {
		p = p.Filter("name", primitive.Regex{Pattern: req.SubAccountName, Options: "i"})
	}

	r.log.Info(ctx, "called %v %v %v %v", page, size, req, p)

	objs := make([]*entity.SubAccount, 0)
	count, err := r.repo.GetAll(p, &objs)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return objs, count, nil
}

func (r subAccountRepoImpl) FindSubAccounts(ctx context.Context, req repository.FindSubAccountsRequest) (map[entity.SubAccountCode]entity.SubAccount, error) {
	r.log.Info(ctx, "called")

	//var bsonObjectIDs bson.A
	//for _, val := range req.SubAccountCodes {
	//	bsonObjectIDs = append(bsonObjectIDs, val)
	//}

	p := database.NewDefaultParam().
		Page(1).
		Size(int64(len(req.SubAccountCodes))).
		Filter("parent_account.wallet_id", req.WalletID).
		Filter("code", bson.M{"$in": SliceToBsonA(req.SubAccountCodes)}).
		Sort("code", 1)

	results := map[entity.SubAccountCode]entity.SubAccount{}
	_, err := r.repo.GetAllEachItem(p, func(obj entity.SubAccount) {
		results[obj.Code] = obj
	})
	if err != nil {
		r.log.Error(ctx, err.Error())
		return nil, err
	}

	return results, nil
}

func (r subAccountRepoImpl) FindAllSubAccountByName(ctx context.Context, walletID entity.WalletID, names []string) (map[string]entity.SubAccount, error) {
	r.log.Info(ctx, "called")

	joinNames := make([]primitive.Regex, 0)

	for _, s := range names {

		joinNames = append(joinNames, primitive.Regex{Pattern: s, Options: "i"})

		//if i == 0 {
		//	joinNames = s
		//	continue
		//}
		//joinNames = fmt.Sprintf("%s|%s", joinNames, s)
	}

	fmt.Printf(">>>>>>>>%s\n", joinNames)

	p := database.NewDefaultParam().
		Page(1).
		Size(int64(len(names))).
		Filter("parent_account.wallet_id", walletID).
		Filter("name", bson.M{"$in": joinNames}).
		Sort("code", 1)

	results := map[string]entity.SubAccount{}
	_, err := r.repo.GetAllEachItem(p, func(obj entity.SubAccount) {
		results[strings.ToLower(obj.Name)] = obj
	})
	if err != nil {
		r.log.Error(ctx, err.Error())
		return nil, err
	}

	return results, nil
}

func (r subAccountRepoImpl) SaveSubAccounts(ctx context.Context, objs []*entity.SubAccount) error {
	r.log.Info(ctx, "called")

	err := r.repo.InsertMany(objs...)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return err
	}

	return nil
}
