package main

import (
	"fmt"

	api "orka/concurrent-deploy/api"
)

func main() {
	cl := api.NewOrkaApiClient()

	vmConfigName := api.GenerateVmConfigName()
	fmt.Println(vmConfigName)

	res := cl.CreateVmConfig(vmConfigName)
	fmt.Println(res)

	res = cl.PurgeVm(vmConfigName)
	fmt.Println(res)
}
