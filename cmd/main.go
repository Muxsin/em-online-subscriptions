package main

import "effective-mobile/online-subscriptions/internal/server"

func main() {
	app := server.New()
	app.Run()
}
