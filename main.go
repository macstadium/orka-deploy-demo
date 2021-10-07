package main

import (
	"fmt"
	"time"

	api "orka/concurrent-deploy/api"
)

const N = 40

func main() {
	var success, failure int
	c := make(chan int, N)

	cl := api.NewOrkaApiClient()
	vmConfigName := api.GenerateVmConfigName()

	cl.CreateVmConfig(vmConfigName)

	fmt.Println("deploying ...")
	start := time.Now()
	for i := 0; i < N; i++ {
		go func() {
			r, _ := cl.DeployVm(vmConfigName)
			c <- r
		}()
	}
	for i := 0; i < N; i++ {
		if <-c == 200 {
			success++
		} else {
			failure++
		}
	}
	duration := time.Since(start)
	fmt.Printf("Total deploy time: %v\n", duration)
	fmt.Printf("Total successful deployments: %v\n", success)
	fmt.Printf("Total failed deployments: %v\n", failure)

	res := cl.DeleteVm(vmConfigName)
	fmt.Println(res)

	cl.PurgeVm(vmConfigName)
}
