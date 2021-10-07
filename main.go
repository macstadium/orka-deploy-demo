package main

import (
	"fmt"

	api "orka/concurrent-deploy/api"
)

func main() {
	cl := api.NewOrkaApiClient()

	vmConfigName := api.GenerateVmConfigName()
	fmt.Println(vmConfigName)

	vmConfig := cl.CreateVmConfig(vmConfigName)
	fmt.Println(vmConfig)
}
