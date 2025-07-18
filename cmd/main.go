package main

import (
	_ "effective-mobile/online-subscriptions/cmd/docs"
	"effective-mobile/online-subscriptions/internal/server"
)

// @title Subscription Service API
// @version 1.0
// @description This is a sample subscription service API.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	app := server.New()
	app.Run()
}
