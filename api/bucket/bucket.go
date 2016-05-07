/*------------------------
name        bucket
describe    api bucket
author      ailn(z.ailn@wmntec.com)
create      2016-05-07
version     1.0
------------------------*/
package bucket

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

//create one bucket
var Post = func(c *gin.Context) {
	//TODO CREATE BUCKET
}

//update one bucket
var Put = func(c *gin.Context) {
	//TODO UPDATE BUCKET
}

//list object in bucket
var Get = func(c *gin.Context) {
	//TODO LIST ALL OBJECT IN BUCKET
}
