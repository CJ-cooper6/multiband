package main

import (
	"multiband/model"
	"multiband/routes"
)

func main() {
	model.IntDb()
	routes.InitRouter()
}
