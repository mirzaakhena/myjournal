package getallsubaccountbalance

import "myjournal/domain_myjournal/model/repository"

// Outport of usecase
type Outport interface {
	repository.FindAllSubAccountBalanceRepo2

	//database.GetAllRepo[entity.SubAccountBalance]
}
