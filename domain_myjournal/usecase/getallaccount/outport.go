package getallaccount

import (
	"myjournal/domain_myjournal/model/repository"
)

// Outport of usecase
type Outport interface {
	repository.FindAllAccountRepo2

	//database.GetAllRepo[entity.Account]
}
