/*------------------------
name        api
describe    api library
author      ailn(z.ailn@wmntec.com)
create      2016-05-05
version     1.0
------------------------*/
package api

import (
	//golang official package
	"fmt"
	"time"

	//third party package
	"github.com/gin-gonic/gin"

	//project package
	"github.com/ailncode/gorgw/api/version"
	"github.com/ailncode/gorgw/base"
	. "github.com/ailncode/gorgw/config"
)

//type of Action
type Action func(c *gin.Context)

//api struct
type Api struct {
	Listen string
}

//method of Api to start up api
func (a *Api) Run() {
	fmt.Println(time.Now().String(), "Run Api with Config:")
	fmt.Println("------------------config------------------")
	for k, v := range Conf {
		fmt.Println(k, "\t\t\t\t\t:\t\t\t", v)
	}
	fmt.Println("------------------------------------------")
	if Conf["debug"] != "true" {
		//switch gin mode to release
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	//use logger middle ware
	router.Use(base.Logger())
	router.GET("/", version.Version)
	authorized := router.Group("/")
	authorized.Use(base.Authorizer())
	{
		//TODO
	}
	router.Run(a.Listen)
}
