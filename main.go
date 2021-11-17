package main

import (
	"fmt"
	"time"
  "flag"

	api "orka/concurrent-deploy/api"
)

func main() {
  numDeploys := flag.Int("deploy", 4, "number of virtual machines to deploy")
  baseImage := flag.String("base-image", "90GBigSurSSH.img", "base image to use for virtual machine configuration")
  cpuCount := flag.Int("cpu-count", 3, "cpu core count to use for virtual machine configuration")
  flag.Parse()

	var success, failure int
	c := make(chan int, *numDeploys)

	cl := api.NewOrkaApiClient()
	vmConfigName := api.GenerateVmConfigName()

	cl.CreateVmConfig(vmConfigName, *baseImage, *cpuCount)

	fmt.Println("deploying ...")
	start := time.Now()
	for i := 0; i < *numDeploys; i++ {
		go func() {
			r, _ := cl.DeployVm(vmConfigName)
			c <- r
		}()
	}
	for i := 0; i < *numDeploys; i++ {
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

	cl.DeleteVm(vmConfigName)
	cl.PurgeVm(vmConfigName)
}
