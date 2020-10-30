package utility

import (
	"encoding/json"
	"os"
)

var Configuration EntConfiguration

type EntConfiguration struct {
	DatabaseDriver   string `json:"dbDriver"`
	DatabaseUserName string `json:"dbUser"`
	DatabasePassword string `json:"dbPass"`
	DatabaseHostName string `json:"hostName"`
	DatabasePort     string `json:"port"`
	DatabaseName     string `json:"dbName"`
}

func InitializeConfig() {
	file, err := os.Open("./env.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Configuration)
	if err != nil {
		panic(err)
	}
}
