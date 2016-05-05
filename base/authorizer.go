/*------------------------
name        authorizer
describe    authorizer middle ware for gin
author      ailn(z.ailn@wmntec.com)
create      2016-05-05
version     1.0
------------------------*/
package base

import (
	//golang official package

	//third party package
	"github.com/gin-gonic/gin"
)

func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO
		c.Next()
	}
}
