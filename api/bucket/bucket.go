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
	"fmt"
	"net/http"

	//third party package
	"github.com/ailncode/golib/mongo"
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"

	//project package
	"github.com/ailncode/gorgw/base"
	. "github.com/ailncode/gorgw/config"
	"github.com/ailncode/gorgw/entity"
)

//create one bucket
var Post = func(c *gin.Context) {
	//TODO CREATE BUCKET
	bucket_name := c.PostForm("bucketname")
	if bucket_name == "" {
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "can not find param bucketname."})
		c.Abort()
		return
	}
	is_public := false
	if c.PostForm("ispublic") == "true" {
		is_public = true
	}
	user := c.MustGet("user").(*entity.User)
	//check
	mgo, err := mongo.NewMongo(Conf["server"])
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, base.ApiErr{http.StatusInternalServerError, "open mongodb server error."})
		c.Abort()
		return
	}
	defer mgo.Close()
	if mgo.IsExist(Conf["db"], Conf["bucketcoll"], map[string]interface{}{"name": bucket_name, "owner": user.Guid}) {
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "bucket name is exist."})
		c.Abort()
		return
	}
	//create bucket
	err = mgo.Insert(Conf["db"], Conf["bucketcoll"], &entity.Bucket{uuid.New(), bucket_name, user.Guid, is_public})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, base.ApiErr{http.StatusInternalServerError, "create bucket server error."})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, base.ApiErr{http.StatusOK, "create bucket success."})
	c.Abort()
	return
}

//update one bucket
var Put = func(c *gin.Context) {
	//TODO UPDATE BUCKET
}

//list object in bucket
var Get = func(c *gin.Context) {
	//TODO LIST ALL OBJECT IN BUCKET
}
