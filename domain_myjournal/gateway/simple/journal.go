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

type journalRepoImpl struct {
	log  logger.Logger
	repo database.Repository[entity.Journal]
}

func NewJournalRepoImpl(log logger.Logger, db *mongo.Database) *journalRepoImpl {
	return &journalRepoImpl{
		log:  log,
		repo: database.NewMongoGateway[entity.Journal](db),
	}
}

func (r journalRepoImpl) FindAllJournal(ctx context.Context, page, size int64, req repository.FindAllJournalRequest) ([]*entity.Journal, int64, error) {
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
		Filter("wallet_id", req.WalletID).
		Filter("user_id", req.UserID).
		Sort("date", -1)

	if len(req.Description) > 0 {
		p = p.Filter("description", primitive.Regex{Pattern: req.Description, Options: "i"})
	}

	if !dateStart.IsZero() {
		p = p.Filter("date", bson.M{"$gte": dateStart})
	}

	if !dateEnd.IsZero() {
		p = p.Filter("date", bson.M{"$lte": dateEnd.AddDate(0, 0, 1)})
	}

	results := make([]*entity.Journal, 0)
	count, err := r.repo.GetAll(p, &results)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return results, count, nil
}

func (r journalRepoImpl) SaveJournal(ctx context.Context, obj *entity.Journal) error {
	r.log.Info(ctx, "called")

	err := r.repo.InsertOrUpdate(obj)
	if err != nil {
		r.log.Error(ctx, err.Error())
		return err
	}

	return nil
}

func (r journalRepoImpl) FindAllJournalByIDs(ctx context.Context, walletId entity.WalletID, journalIDs []entity.JournalID) (map[entity.JournalID]entity.Journal, error) {
	r.log.Info(ctx, "called")

	p := database.NewDefaultParam().
		Page(1).
		Size(200).
		Filter("wallet_id", walletId).
		Sort("date", -1)

	result := map[entity.JournalID]entity.Journal{}
	_, err := r.repo.GetAllEachItem(p, func(obj entity.Journal) {
		result[obj.ID] = obj
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
