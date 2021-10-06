package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	conf "orka/concurrent-deploy/conf"
)

type VmConfig struct {
	VmName    string `json:"orka_vm_name"`
	BaseImage string `json:"orka_base_image"`
	CPU       int    `json:"orka_cpu_core"`
}

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

func CreateVmConfig(oc conf.OrkaConf, vmConfigName string) string {
	postBody, _ := json.Marshal(VmConfig{VmName: vmConfigName, BaseImage: "90GBigSurSSH.img", CPU: 3})
	req, err := http.NewRequest(http.MethodPost, oc.URL+"/resources/vm/create", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oc.Token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}
