package main

import (
	"fmt"
	"time"

	api "orka/concurrent-deploy/api"
)

const N = 3

func main() {
	c := make(chan string, N)
	cl := api.NewOrkaApiClient()

	vmConfigName := api.GenerateVmConfigName()
	fmt.Println(vmConfigName)

	cl.CreateVmConfig(vmConfigName)

	start := time.Now()
	for i := 0; i < N; i++ {
		go func() {
			r := cl.DeployVm(vmConfigName)
			c <- r
		}()
	}
	for i := 0; i < N; i++ {
		fmt.Println(<-c)
	}
	duration := time.Since(start)
	fmt.Printf("Total deploy time: %v\n", duration)

	cl.PurgeVm(vmConfigName)
}
