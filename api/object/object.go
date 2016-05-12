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
	"mime/multipart"
	"net/http"

	//third party package
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"

	//project package
	"github.com/ailncode/gorgw/base"
	"github.com/ailncode/gorgw/entity"
	"github.com/ailncode/gorgw/lib/bucket"
	"github.com/ailncode/gorgw/lib/object"
)

//create one object
var Post = func(c *gin.Context) {
	r, err := c.Request.MultipartReader()
	if err != nil {
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "bad multipart data."})
		c.Abort()
		return
	}
	form := make(map[string]string)
	file := make(map[string]*multipart.Part)
	for {
		p, err := r.NextPart()
		if err != nil {
			break
		}
		if p.FileName() != "" {
			file[p.FormName()] = p
		} else {
			defer p.Close()
			b, err := ioutil.ReadAll(p)
			if err != nil {
				fmt.Println(err)
				continue
			}
			form[p.FormName()] = string(b)
		}
	}

	key, ok := form["key"]
	if !ok {
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "can not find param key."})
		c.Abort()
		return
	}
	md5, ok := form["md5"]
	if !ok {
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "can not find param md5."})
		c.Abort()
		return
	}
	f, ok := file["file"]
	if !ok {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, base.ApiErr{http.StatusBadRequest, "can not find param file."})
		c.Abort()
		return
	}
	//write object in ceph
	user := c.MustGet("user").(*entity.User)
	bucket_name := c.Param("bucketname")
	b, err := bucket.Get(user.Guid, bucket_name)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, base.ApiErr{http.StatusInternalServerError, "get bucket error."})
		c.Abort()
		return
	}
	task_id := uuid.New()
	//TODO CREATE TASK
	go func() {
		object.Create(b.Guid, key, md5, f, task_id)
	}()
	c.JSON(http.StatusOK, entity.Operate{"Create Object", task_id})
	c.Abort()
	return
}

//update one object
var Put = func(c *gin.Context) {

}

//get one object
var Get = func(c *gin.Context) {

}
