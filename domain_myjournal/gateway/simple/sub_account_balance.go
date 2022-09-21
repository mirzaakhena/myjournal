package simple

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"myjournal/domain_myjournal/model/entity"
	"myjournal/domain_myjournal/model/repository"
	"myjournal/shared/infrastructure/database"
	"myjournal/shared/infrastructure/logger"
)

type subAccountBalanceRepoImpl struct {
	log  logger.Logger
	repo database.Repository[entity.SubAccountBalance]
	DB   *mongo.Database
}

func NewSubAccountBalanceRepoImpl(log logger.Logger, db *mongo.Database) *subAccountBalanceRepoImpl {
	return &subAccountBalanceRepoImpl{
		log:  log,
		repo: database.NewMongoGateway[entity.SubAccountBalance](db),
		DB:   db,
	}
}

func (r subAccountBalanceRepoImpl) SaveSubAccountBalances(ctx context.Context, objs []*entity.SubAccountBalance) error {
	r.log.Info(ctx, "called")

	err := r.repo.InsertMany(objs...)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return err
	}

	return nil
}

func (r subAccountBalanceRepoImpl) FindLastSubAccountBalances(ctx context.Context, req repository.FindLastSubAccountBalancesRequest) (map[entity.SubAccountCode]entity.Money, error) {
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

	coll := r.DB.Collection(r.repo.GetTypeName())
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

	type codeBalance struct {
		ID      entity.SubAccountCode `bson:"_id"`
		Balance entity.Money          `bson:"balance"`
	}

	for cursor.Next(ctx) {

		var result codeBalance
		if err := cursor.Decode(&result); err != nil {
			r.log.Error(ctx, err.Error())
			return nil, err
		}

		results[result.ID] = result.Balance
	}
	if err := cursor.Err(); err != nil {
		r.log.Error(ctx, err.Error())
		return nil, err
	}

	return results, nil
}

func (r subAccountBalanceRepoImpl) FindAllSubAccountBalance(ctx context.Context, page, size int64, req repository.FindAllSubAccountBalanceRequest) ([]*entity.SubAccountBalance, int64, error) {
	r.log.Info(ctx, "called")

	dateStart, err := req.DateStart.GetTime()
	if err != nil {
		return nil, 0, err
	}

	dateEnd, err := req.DateEnd.GetTime()
	if err != nil {
		return nil, 0, err
	}

	p := database.NewDefaultParam().
		Page(page).
		Size(size).
		Filter("sub_account.parent_account.wallet_id", req.WalletID).
		Sort("date", -1)

	if len(req.SubAccountName) > 0 {
		p = p.Filter("sub_account.name", primitive.Regex{Pattern: req.SubAccountName, Options: "i"})
	}

	if !dateStart.IsZero() {
		p = p.Filter("date", bson.M{"$gte": dateStart})
	}

	if !dateEnd.IsZero() {
		p = p.Filter("date", bson.M{"$lte": dateEnd.AddDate(0, 0, 1)})
	}

	results := make([]*entity.SubAccountBalance, 0)
	count, err := r.repo.GetAll(p, &results)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return results, count, nil
}

func (r subAccountBalanceRepoImpl) FindAllSubAccountBalanceByJournalID(ctx context.Context, walletId entity.WalletID, journalIDs []entity.JournalID) (map[entity.JournalID][]entity.SubAccountBalance, error) {
	r.log.Info(ctx, "called")

	var bsonObjectIDs bson.A
	for _, val := range journalIDs {
		bsonObjectIDs = append(bsonObjectIDs, val)
	}

	p := database.NewDefaultParam().
		Page(1).
		Size(200).
		Filter("sub_account.parent_account.wallet_id", walletId).
		Filter("journal_id", bson.M{"$in": bsonObjectIDs}).
		Sort("sequence", 1)

	result := map[entity.JournalID][]entity.SubAccountBalance{}
	_, err := r.repo.GetAllEachItem(p, func(obj entity.SubAccountBalance) {
		if _, exist := result[obj.JournalID]; !exist {
			result[obj.JournalID] = make([]entity.SubAccountBalance, 0)
		}
		result[obj.JournalID] = append(result[obj.JournalID], obj)
	})
	if err != nil {
		r.log.Error(ctx, err.Error())
		return nil, err
	}

	return result, nil
}
