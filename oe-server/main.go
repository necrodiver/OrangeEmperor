package main

import (
	"oe_server/helpers/utils"
	route "oe_server/routers"
)

func main() {
	utils.InitZap()
	route.InitRouter()
}
