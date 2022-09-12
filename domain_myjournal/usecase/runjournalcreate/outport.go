package runjournalcreate

import "myjournal/domain_myjournal/model/repository"

// Outport of usecase
type Outport interface {
	repository.SaveJournalRepo
	repository.SaveSubAccountBalancesRepo
	repository.SaveAccountBalancesRepo
	repository.FindLastSubAccountBalancesRepo
	repository.FindSubAccountsRepo
}
