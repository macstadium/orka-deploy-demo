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

type OrkaApiClient struct {
  Client *http.Client
  Conf conf.OrkaConf
}

func NewOrkaApiClient() *OrkaApiClient {
  return &OrkaApiClient{
    Client: &http.Client{},
    Conf: conf.ReadConf(),
  }
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

func (cl *OrkaApiClient) CreateVmConfig(vmConfigName string) string {
	postBody, _ := json.Marshal(map[string]interface{}{"orka_vm_name": vmConfigName, "orka_base_image": "90GBigSurSSH.img", "orka_cpu_core": 3})
	req, err := http.NewRequest(http.MethodPost, cl.Conf.URL+"/resources/vm/create", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.Conf.Token))

	res, err := cl.Client.Do(req)
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
