package restapi

import (
	"github.com/gin-gonic/gin"

	"myjournal/domain_myjournal/usecase/getallaccountbalance"
	"myjournal/domain_myjournal/usecase/getalljournal"
	"myjournal/domain_myjournal/usecase/getallsubaccountbalance"
	"myjournal/domain_myjournal/usecase/runaccountscreate"
	"myjournal/domain_myjournal/usecase/runjournalcreate"
	"myjournal/domain_myjournal/usecase/runjournalrollback"
	"myjournal/domain_myjournal/usecase/runsubaccountscreate"
	"myjournal/domain_myjournal/usecase/runwalletcreate"
	"myjournal/shared/infrastructure/config"
	"myjournal/shared/infrastructure/logger"
)

type Controller struct {
	Router                        gin.IRouter
	Config                        *config.Config
	Log                           logger.Logger
	GetAllAccountBalanceInport    getallaccountbalance.Inport
	GetAllJournalInport           getalljournal.Inport
	GetAllSubaccountBalanceInport getallsubaccountbalance.Inport
	RunAccountsCreateInport       runaccountscreate.Inport
	RunJournalCreateInport        runjournalcreate.Inport
	RunJournalRollbackInport      runjournalrollback.Inport
	RunSubAccountsCreateInport    runsubaccountscreate.Inport
	RunWalletCreateInport         runwalletcreate.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("/getallaccountbalance", r.authorized(), r.getAllAccountBalanceHandler())
	r.Router.GET("/getalljournal", r.authorized(), r.getAllJournalHandler())
	r.Router.GET("/getallsubaccountbalance", r.authorized(), r.getAllSubaccountBalanceHandler())

	r.Router.POST("/accounts", r.authorized(), r.runAccountsCreateHandler())
	r.Router.POST("/subaccounts", r.authorized(), r.runSubAccountsCreateHandler())

	r.Router.POST("/wallet/:walletId/journal", r.authorized(), r.runJournalCreateHandler())
	r.Router.POST("/runwalletcreate", r.authorized(), r.runWalletCreateHandler())
	r.Router.POST("/runjournalrollback", r.authorized(), r.runJournalRollbackHandler())
}
