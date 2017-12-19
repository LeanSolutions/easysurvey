package main

import (
	"EasySurvey/app/shared/server"
	"encoding/json"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

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
