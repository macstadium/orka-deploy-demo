package main

import (
	"fmt"
	"time"

	api "orka/concurrent-deploy/api"
)

func main() {
	cl := api.NewOrkaApiClient()

	vmConfigName := api.GenerateVmConfigName()
	fmt.Println(vmConfigName)

	res := cl.CreateVmConfig(vmConfigName)
	fmt.Println(res)

	start := time.Now()
	res = cl.DeployVm(vmConfigName)
	fmt.Println(res)
	duration := time.Since(start)
	fmt.Printf("Total deploy time: %v\n", duration)

	res = cl.PurgeVm(vmConfigName)
	fmt.Println(res)
}
