package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/MikelSot/metal-bat/model"
)

func newConfiguration(path string) model.Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	conf := model.Configuration{}
	if err := json.Unmarshal(file, &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}
