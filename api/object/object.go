/*------------------------
name        object
describe    api object
author      ailn(z.ailn@wmntec.com)
create      2016-05-07
version     1.0
------------------------*/
package object

import (
	//golang official package
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	//third party package
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"

	//project package
	"github.com/ailncode/gorgw/base"
	"github.com/ailncode/gorgw/entity"
	"github.com/ailncode/gorgw/lib/bucket"
	"github.com/ailncode/gorgw/lib/object"
	"github.com/ailncode/gorgw/lib/task"
)

//create one object
var Post = func(c *gin.Context) {

	r, err := c.Request.MultipartReader()
	if err != nil {
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "bad multipart data."})
		c.Abort()
		return
	}
	p, err := r.NextPart()
	if err != nil {
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "can not find first param key."})
		c.Abort()
		return
	}
	defer p.Close()
	key_buff, err := ioutil.ReadAll(p)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, base.ApiErr{http.StatusInternalServerError, "read first param key error."})
		c.Abort()
		return
	}
	key := string(key_buff)
	if key == "" {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "param key must be non null string."})
		c.Abort()
		return
	}
	bucket_name := c.Param("bucketname")
	user := c.MustGet("user").(*entity.User)
	b, err := bucket.Get(user.Guid, bucket_name)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, base.ApiErr{http.StatusInternalServerError, "get bucket error."})
		c.Abort()
		return
	}
	exist, err := object.IsExist(b.Guid, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, base.ApiErr{http.StatusInternalServerError, "check object isexist in this bucket error."})
		c.Abort()
		return
	}
	if exist {
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "object is exist in this bucket."})
		c.Abort()
		return
	}
	p, err = r.NextPart()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "can not find second param md5."})
		c.Abort()
		return
	}
	defer p.Close()
	md5_buff, err := ioutil.ReadAll(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "read second param md5 error."})
		c.Abort()
		return
	}
	md5 := string(md5_buff)
	if md5 == "" {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "param md5 must be non null string."})
		c.Abort()
		return
	}
	p, err = r.NextPart()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "can not find second param file."})
		c.Abort()
		return
	}
	defer p.Close()
	task_id := uuid.New()
	//TODO CREATE TASK
	err = task.Create(&entity.Task{task_id, time.Now().Unix(), "create object " + key, 0, ""})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, base.ApiErr{http.StatusInternalServerError, "can not create task."})
		c.Abort()
		return
	}
	err = object.Create(b.Guid, key, bucket_name, md5, p, task_id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, base.ApiErr{http.StatusInternalServerError, "create object error."})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, base.ApiErr{http.StatusOK, "create object success."})
	c.Abort()
	return
}

//update one object
var Put = func(c *gin.Context) {

}

//get one object
var Get = func(c *gin.Context) {

}
