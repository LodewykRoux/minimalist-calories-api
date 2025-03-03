package main

import (
	"minimalist-calories-api/initializers"
	"minimalist-calories-api/routing"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := routing.SetupRouter()

	r.Run()
}
