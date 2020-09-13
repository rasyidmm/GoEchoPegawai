package config

import "github.com/tkanos/gonfig"

type Conguration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT 	string
	DB_HOST 	string
	DB_NAME 	string
}

func GetConfig() Conguration {
	conf := Conguration{}
	gonfig.GetConf("config/config.json",&conf)
	return conf

}
