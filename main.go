package main

import (
	"fmt"
	cli "orka/concurrent-deploy/cli"
)

func main() {
	orkaConf := cli.ReadConf()

	fmt.Println(orkaConf)
}
