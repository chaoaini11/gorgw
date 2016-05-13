/*------------------------
name        check bucket
describe    check bucket middle ware for gin
author      ailn(z.ailn@wmntec.com)
create      2016-05-11
version     1.0
------------------------*/
package base

import (
	//golang official package
	"net/http"

	//third party package
	"github.com/gin-gonic/gin"

	//project package
	"github.com/ailncode/gorgw/entity"
	"github.com/ailncode/gorgw/lib/bucket"
)

func CheckBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*entity.User)
		bucket_name := c.Param("bucketname")
		isexist, err := bucket.IsExist(user.Guid, bucket_name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ApiErr{http.StatusInternalServerError, err.Error()})
			c.Abort()
			return
		}
		if !isexist {
			c.JSON(http.StatusForbidden, ApiErr{http.StatusForbidden, "can not find this bucket in you account."})
			c.Abort()
			return
		}
		c.Next()
	}
}
