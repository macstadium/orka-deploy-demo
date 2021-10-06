package main

import (
	"fmt"

	api "orka/concurrent-deploy/api"
	conf "orka/concurrent-deploy/conf"
)

func main() {
	orkaConf := conf.ReadConf()

	fmt.Println(orkaConf)

	healthCheck := api.HealthCheck(orkaConf)
	fmt.Println(healthCheck)

	vmConfigName := api.GenerateVmConfigName()
	fmt.Println(vmConfigName)
}
