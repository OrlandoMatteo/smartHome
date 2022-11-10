package main

import (
	"github.com/OrlandoMatteo/smartHome/lib"
)

func main() {
	conf := lib.GetConfig("config.json")
	lib.CreateCompose(*conf)
}
