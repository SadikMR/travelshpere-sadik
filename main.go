package main

import (
	_ "travelshpere-sadik/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

