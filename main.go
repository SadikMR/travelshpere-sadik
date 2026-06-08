package main

import (
	_ "github.com/SadikMR/travelshpere-sadik/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
