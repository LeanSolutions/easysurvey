package main

import (
	"github.com/lavpthak/easysurvey/app/route"
	"github.com/lavpthak/easysurvey/app/shared/jsonconfig"
	"github.com/lavpthak/easysurvey/app/shared/server"
	"encoding/json"
	"os"
	"runtime"
	"github.com/lavpthak/easysurvey/app/controller"
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
