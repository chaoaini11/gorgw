/*------------------------
name        main
describe    api entrance
author      ailn(z.ailn@wmntec.com)
create      2016-05-05
version     1.0
------------------------*/
package main

import (
	//project package
	"github.com/ailncode/gorgw/api"
	. "github.com/ailncode/gorgw/config"
)

//func main
func main() {
	LoadConfig("config/gorgw.conf")
	a := api.Api{Conf["listen"] + ":" + Conf["port"]}
	a.Run()
}
