package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"

	_ "tt-dev-test/env"
	_ "tt-dev-test/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
