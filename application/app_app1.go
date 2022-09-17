package application

import (
	"myjournal/domain_myjournal/controller/restapi"
	"myjournal/domain_myjournal/gateway/newprod"
	"myjournal/domain_myjournal/usecase/getallaccount"
	"myjournal/domain_myjournal/usecase/getallaccountbalance"
	"myjournal/domain_myjournal/usecase/getalljournal"
	"myjournal/domain_myjournal/usecase/getallsubaccount"
	"myjournal/domain_myjournal/usecase/getallsubaccountbalance"
	"myjournal/domain_myjournal/usecase/runaccountscreate"
	"myjournal/domain_myjournal/usecase/runjournalcreate"
	"myjournal/domain_myjournal/usecase/runsubaccountscreate"
	"myjournal/domain_myjournal/usecase/runwalletcreate"
	"myjournal/shared/driver"
	"myjournal/shared/infrastructure/config"
	"myjournal/shared/infrastructure/logger"
	"myjournal/shared/infrastructure/server"
	"myjournal/shared/util"
)

type app1 struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c app1) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewApp1() func() driver.RegistryContract {

	const appName = "app1"

	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData(appName, appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

		datasource := newprod.NewGateway(log, appData, cfg)

		return &app1{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Router:                        httpHandler.Router,
				Config:                        cfg,
				Log:                           log,
				GetAllAccountBalanceInport:    getallaccountbalance.NewUsecase(datasource),
				GetAllJournalInport:           getalljournal.NewUsecase(datasource),
				GetAllSubaccountBalanceInport: getallsubaccountbalance.NewUsecase(datasource),
				RunAccountsCreateInport:       runaccountscreate.NewUsecase(datasource),
				RunJournalCreateInport:        runjournalcreate.NewUsecase(datasource),
				RunSubAccountsCreateInport:    runsubaccountscreate.NewUsecase(datasource),
				RunWalletCreateInport:         runwalletcreate.NewUsecase(datasource),
				GetAllAccountInport:           getallaccount.NewUsecase(datasource),
				GetAllSubAccountInport:        getallsubaccount.NewUsecase(datasource),
			},
		}

	}
}
