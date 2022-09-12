package runaccountscreate

import (
	"myjournal/domain_myjournal/model/repository"
)

// Outport of usecase
type Outport interface {
	repository.SaveAccountsRepo
}
