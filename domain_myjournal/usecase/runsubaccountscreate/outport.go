package runsubaccountscreate

import "myjournal/domain_myjournal/model/repository"

// Outport of usecase
type Outport interface {
	repository.SaveSubAccountsRepo
	repository.FindAccountsRepo

	//database.GetAllRepo[entity.Account]
	//database.InsertManyRepo[entity.SubAccount]
	//database.InsertManyRepo[entity.Account]
}
