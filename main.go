package main

import routes2 "trolley-backend/routes"

func main() {
	routes := routes2.Init()
	routes.Logger.Fatal(routes.Start(":8030"))
}