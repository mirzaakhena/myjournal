package newprod

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	AccountRepo           database.Repository[entity.Account]
	SubAccountRepo        database.Repository[entity.SubAccount]
	SubAccountBalanceRepo database.Repository[entity.SubAccountBalance]
	JournalRepo           database.Repository[entity.Journal]

	Database *mongo.Database
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	db := database.NewDatabase()

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,

		Database: db,

		AccountRepo:           &database.MongoGateway[entity.Account]{Database: db},
		SubAccountRepo:        &database.MongoGateway[entity.SubAccount]{Database: db},
		SubAccountBalanceRepo: &database.MongoGateway[entity.SubAccountBalance]{Database: db},
		JournalRepo:           &database.MongoGateway[entity.Journal]{Database: db},
	}
}

func (r *gateway) FindAllAccount(ctx context.Context) database.GetAllRepo[entity.Account] {
	r.log.Info(ctx, "called")
	return r.AccountRepo
}

func (r *gateway) FindAllJournal(ctx context.Context) database.GetAllRepo[entity.Journal] {
	r.log.Info(ctx, "called")
	return r.JournalRepo
}

func (r *gateway) FindAllSubAccount(ctx context.Context) database.GetAllRepo[entity.SubAccount] {
	r.log.Info(ctx, "called")
	return r.SubAccountRepo
}

func (r *gateway) FindAllSubAccountBalance(ctx context.Context) database.GetAllRepo[entity.SubAccountBalance] {
	r.log.Info(ctx, "called")
	return r.SubAccountBalanceRepo
}

func (r *gateway) SaveAccounts(ctx context.Context) database.InsertManyRepo[entity.Account] {
	r.log.Info(ctx, "called")
	return r.AccountRepo
}

func (r *gateway) SaveJournal(ctx context.Context) database.InsertOrUpdateRepo[entity.Journal] {
	r.log.Info(ctx, "called")
	return r.JournalRepo
}

func (r *gateway) SaveSubAccountBalances(ctx context.Context) database.InsertManyRepo[entity.SubAccountBalance] {
	r.log.Info(ctx, "called")
	return r.SubAccountBalanceRepo
}

func (r *gateway) FindLastSubAccountBalances(ctx context.Context, req repository.FindLastSubAccountBalancesRequest) (map[entity.SubAccountCode]entity.Money, error) {
	r.log.Info(ctx, "called")

	var subAccountCodes bson.A
	for _, val := range req.SubAccountCodes {
		subAccountCodes = append(subAccountCodes, val)
	}

	//---
	matchStage := bson.D{
		{
			"$match",
			bson.M{
				"sub_account.parent_account.wallet_id": req.WalletID,
				"sub_account.code":                     bson.M{"$in": subAccountCodes},
			},
		},
	}

	groupStage := bson.D{
		{
			"$group",
			bson.M{
				"_id":     "$sub_account.code",
				"balance": bson.M{"$last": "$balance"},
			},
		},
	}

	sortStage := bson.D{
		{
			"$sort",
			bson.M{"date": -1},
		},
	}
	//---

	coll := r.Database.Collection("sub_account_balance")
	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{
		matchStage,
		groupStage,
		sortStage,
	})
	if err != nil {
		r.log.Error(ctx, err.Error())
		return nil, err
	}

	results := map[entity.SubAccountCode]entity.Money{}

	for cursor.Next(ctx) {

		var result bson.D
		if err := cursor.Decode(&result); err != nil {
			r.log.Error(ctx, err.Error())
			return nil, err
		}

		x := result.Map()
		codeAsAny := x["_id"]
		balanceAsAny := x["balance"]

		codeAsString, ok := codeAsAny.(string)
		if !ok {
			panic(err.Error())
		}

		balanceAsFloat64, ok := balanceAsAny.(float64)
		if !ok {
			panic(err.Error())
		}

		results[entity.SubAccountCode(codeAsString)] = entity.Money(balanceAsFloat64)

		//results[result.] = &result
		//r.log.Info(ctx, "%+v", result)
	}
	if err := cursor.Err(); err != nil {
		r.log.Error(ctx, err.Error())
		return nil, err
	}

	return results, nil
}

func (r *gateway) FindSubAccounts(ctx context.Context) database.GetAllEachItemRepo[entity.SubAccount] {
	r.log.Info(ctx, "called")
	return r.SubAccountRepo
}

func (r *gateway) SaveSubAccounts(ctx context.Context) database.InsertManyRepo[entity.SubAccount] {
	r.log.Info(ctx, "called")
	return r.SubAccountRepo
}

func (r *gateway) FindAccounts(ctx context.Context) database.GetAllEachItemRepo[entity.Account] {
	r.log.Info(ctx, "called")
	return r.AccountRepo
}
