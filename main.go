package main

import (
	"flag"
	"fmt"
	"myjournal/application"
	"myjournal/shared/driver"
)

var Version = "0.0.1"

func main() {
	appMap := map[string]func() driver.RegistryContract{
		"app1": application.NewApp1(),
	}
	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if exist {
		fmt.Printf("Version %s", Version)
		driver.Run(app())
	} else {
		fmt.Println("You may try 'go run main.go <app_name>' :")
		for appName := range appMap {
			fmt.Printf(" - %s\n", appName)
		}
	}

}
