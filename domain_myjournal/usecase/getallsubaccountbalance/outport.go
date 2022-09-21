package getallsubaccountbalance

import "myjournal/domain_myjournal/model/repository"

// Outport of usecase
type Outport interface {
	repository.FindAllSubAccountBalanceRepo
	repository.

		//database.GetAllRepo[entity.SubAccountBalance]
		FindAllJournalByIDsRepo
}
