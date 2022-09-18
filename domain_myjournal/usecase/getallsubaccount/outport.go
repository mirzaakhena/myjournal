package getallsubaccount

import (
	"myjournal/domain_myjournal/model/repository"
)

// Outport of usecase
type Outport interface {
	repository.FindAllSubAccountRepo2

	//database.GetAllRepo[entity.SubAccount]
}
