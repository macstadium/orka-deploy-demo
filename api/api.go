package api

import (
	"io/ioutil"
	"log"
	"net/http"

	conf "orka/concurrent-deploy/conf"
)

func HealthCheck(oc conf.OrkaConf) string {
	r, err := http.Get(oc.URL + "/health-check")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}
