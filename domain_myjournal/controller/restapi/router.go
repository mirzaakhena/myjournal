package restapi

import (
	"github.com/gin-gonic/gin"

	"myjournal/domain_myjournal/usecase/getallaccount"
	"myjournal/domain_myjournal/usecase/getallaccountbalance"
	"myjournal/domain_myjournal/usecase/getalljournal"
	"myjournal/domain_myjournal/usecase/getallsubaccount"
	"myjournal/domain_myjournal/usecase/getallsubaccountbalance"
	"myjournal/domain_myjournal/usecase/runaccountscreate"
	"myjournal/domain_myjournal/usecase/runjournalcreate"
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
	RunSubAccountsCreateInport    runsubaccountscreate.Inport
	RunWalletCreateInport         runwalletcreate.Inport
	GetAllAccountInport           getallaccount.Inport
	GetAllSubAccountInport        getallsubaccount.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("/wallet/:walletId/journal", r.authorized(), r.getAllJournalHandler())
	r.Router.GET("/wallet/:walletId/accountbalance", r.authorized(), r.getAllAccountBalanceHandler())
	r.Router.GET("/wallet/:walletId/subaccountbalance", r.authorized(), r.getAllSubaccountBalanceHandler())

	r.Router.POST("/wallet/:walletId/accounts", r.authorized(), r.runAccountsCreateHandler())
	r.Router.POST("/wallet/:walletId/subaccounts", r.authorized(), r.runSubAccountsCreateHandler())

	r.Router.POST("/wallet/:walletId/journal", r.authorized(), r.runJournalCreateHandler())
	r.Router.POST("/wallet", r.authorized(), r.runWalletCreateHandler())

	r.Router.GET("/wallet/:walletId/account", r.authorized(), r.getAllAccountHandler())
	r.Router.GET("/wallet/:walletId/subaccount", r.authorized(), r.getAllSubAccountHandler())
}
