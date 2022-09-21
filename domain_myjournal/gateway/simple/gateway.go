package simple

import (
	"myjournal/shared/driver"
	"myjournal/shared/infrastructure/config"
	"myjournal/shared/infrastructure/database"
	"myjournal/shared/infrastructure/logger"
)

type gateway struct {
	*accountRepoImpl
	*journalRepoImpl
	*subAccountRepoImpl
	*subAccountBalanceRepoImpl
	//log     logger.Logger
	//appData driver.ApplicationData
	//config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	db := database.NewDatabase()

	return &gateway{
		accountRepoImpl:           NewAccountRepoImpl(log, db),
		journalRepoImpl:           NewJournalRepoImpl(log, db),
		subAccountRepoImpl:        NewSubAccountRepoImpl(log, db),
		subAccountBalanceRepoImpl: NewSubAccountBalanceRepoImpl(log, db),
		//log:                       log,
		//appData:                   appData,
		//config:                    cfg,
	}
}
