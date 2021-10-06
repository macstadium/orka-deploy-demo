package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type OrkaConf struct {
	URL   string `json:"api-url"`
	Token string `json:"token"`
}

func (oc OrkaConf) String() string {
	b, err := json.Marshal(oc)
	if err != nil {
		log.Fatalln(err)
	}
	return string(b)
}

func ReadConf() OrkaConf {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	orkaConfPath := filepath.Join(dir, ".config", "configstore", "orka-cli.json")

	f, err := ioutil.ReadFile(orkaConfPath)
	if err != nil {
		log.Fatalln(err)
	}

	var orkaConf OrkaConf
	err = json.Unmarshal([]byte(f), &orkaConf)
	if err != nil {
		log.Fatalln(err)
	}
	return orkaConf
}
