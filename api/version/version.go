/*------------------------
name        version
describe    api version
author      ailn(z.ailn@wmntec.com)
create      2016-05-05
version     1.0
------------------------*/
package version

import (
	//golang official package
	"net/http"
	"strconv"
	"time"

	//third party package
	"github.com/gin-gonic/gin"

	//project package
	"github.com/ailncode/gorgw/base"
)

//action of api version
var Version = func(c *gin.Context) {
	c.JSON(http.StatusOK, base.ApiErr{200, strconv.FormatInt(time.Now().Unix(), 10) + " api version 1.0"})
}
