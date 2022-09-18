package runjournalcreate

import "myjournal/domain_myjournal/model/repository"

// Outport of usecase
type Outport interface {
	repository.SaveJournalRepo2
	repository.SaveSubAccountBalancesRepo2
	repository.FindLastSubAccountBalancesRepo
	repository.FindSubAccountsRepo2
	//repository.SaveAccountBalancesRepo2
}
