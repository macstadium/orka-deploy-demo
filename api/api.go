package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	conf "orka/deploy-demo/conf"
)

type OrkaApiClient struct {
	Client *http.Client
	Conf   conf.OrkaConf
}

func NewOrkaApiClient() *OrkaApiClient {
	return &OrkaApiClient{
		Client: &http.Client{Timeout: time.Minute * 1},
		Conf:   conf.ReadConf(),
	}
}

func (cl *OrkaApiClient) CallApi(method string, route string, b []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, cl.Conf.URL+route, bytes.NewBuffer(b))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.Conf.Token))

	return cl.Client.Do(req)
}

func (cl *OrkaApiClient) CreateVmConfig(vmConfigName string, baseImage string, cpuCount int) string {
	reqBody, _ := json.Marshal(map[string]interface{}{"orka_vm_name": vmConfigName, "orka_base_image": baseImage, "orka_cpu_core": cpuCount})
	res, err := cl.CallApi(http.MethodPost, "/resources/vm/create", reqBody)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != http.StatusCreated {
		log.Fatalln("Unable to create VM config!", string(b))
	}

	return string(b)
}

func (cl *OrkaApiClient) DeployVm(vmConfigName string) (int, string) {
	reqBody, _ := json.Marshal(map[string]string{"orka_vm_name": vmConfigName})
	res, err := cl.CallApi(http.MethodPost, "/resources/vm/deploy", reqBody)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return res.StatusCode, string(b)
}

func (cl *OrkaApiClient) DeleteVm(vmConfigName string) {
	reqBody, _ := json.Marshal(map[string]string{"orka_vm_name": vmConfigName})
	res, err := cl.CallApi(http.MethodDelete, "/resources/vm/delete", reqBody)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Unable to delete VM: %v\n", vmConfigName)
	}
}

func (cl *OrkaApiClient) PurgeVm(vmConfigName string) {
	reqBody, _ := json.Marshal(map[string]string{"orka_vm_name": vmConfigName})
	res, err := cl.CallApi(http.MethodDelete, "/resources/vm/purge", reqBody)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Unable to purge VM config: %v\n", vmConfigName)
	}
}
