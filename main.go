package main

import (
	"EasySurvey/app/route"
	"EasySurvey/app/shared/jsonconfig"
	"EasySurvey/app/shared/server"
	"encoding/json"
	"os"
	"runtime"
	"EasySurvey/app/controller"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	jsonconfig.Load("app"+string(os.PathSeparator)+"config"+string(os.PathSeparator)+"config.json", config)

	controller.Load()

	server.Run(route.LoadHTTP(), config.Server)
}

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Server server.Server `json:"Server"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
