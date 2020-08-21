package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	Port   int
	DbUrl  string
	DbName string
}

func GetConfig() (config Config, err error) {

	data, err := ioutil.ReadFile("conf.json")

	if err != nil {
		fmt.Println(err)
	}

	conf := Config{}
	json.Unmarshal(data, &conf)

	if err != nil {
		log.Print(err)
		return Config{}, err
	}
	return conf, nil
}
