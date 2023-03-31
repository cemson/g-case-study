package main

import (
	"fmt"
	"g-case-study/controllers"
	"g-case-study/globals"
	"g-case-study/logging"
	"g-case-study/settings"
	"g-case-study/utilities"
	"net/http"
)

func main() {
	appSettings := settings.AppSettings{}
	// loading application configuration from conf.json file with reflection
	utilities.LoadConfig(&appSettings)
	// set a global appSettings instance for reaching application configuration everywhere
	globals.SetAppSettings(&appSettings)
	initializeHandlers()
	logging.Log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", appSettings.Port), nil))
}
func initializeHandlers() {
	testController := controllers.CreateTestController()
	keyValController := controllers.CreateKeyValValidator()
	recordsController := controllers.CreateRecordsController()

	http.HandleFunc("/test", testController.HandleHealthCheck)

	// mongodb
	http.HandleFunc("/records", recordsController.Handle)

	// in - memory
	http.HandleFunc("/keyval", keyValController.Handle)
}
