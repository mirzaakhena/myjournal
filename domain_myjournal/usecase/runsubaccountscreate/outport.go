package runsubaccountscreate

import "myjournal/domain_myjournal/model/repository"

// Outport of usecase
type Outport interface {
	repository.SaveSubAccountsRepo
	repository.FindAccountsRepo
}
