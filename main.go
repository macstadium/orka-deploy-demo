package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	conf "orka/concurrent-deploy/conf"
)

func main() {
	orkaConf := conf.ReadConf()

	fmt.Println(orkaConf)

	r, err := http.Get(orkaConf.URL + "/health-check")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}
