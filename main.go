package main 

import (
	DB "gymfinity-backend-api/Connection"
	"gymfinity-backend-api/Router"
)

func main() {
	DB.Connect();
	Router.SetupRoutes();
}