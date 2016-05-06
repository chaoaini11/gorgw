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
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"math"
	"net/http"
	"sort"
	"strings"
	"time"

	//third party package
	"github.com/gin-gonic/gin"
)

func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		date := c.Request.Header.Get("Date")
		if date == "" {
			c.JSON(http.StatusBadRequest, ApiErr{http.StatusBadRequest, "can not find HTTP Date Header."})
			c.Abort()
			return
		}
		authorizer := c.Request.Header.Get("Authorization")
		if authorizer == "" {
			c.JSON(http.StatusBadRequest, ApiErr{http.StatusBadRequest, "can not find HTTP Authorization Header."})
			c.Abort()
			return
		}
		id_key := strings.Split(authorizer, ":")
		if len(id_key) != 2 {
			c.JSON(http.StatusBadRequest, ApiErr{http.StatusBadRequest, "bad signature."})
			c.Abort()
			return
		}
		//parse http date
		var t time.Time
		if t, err = time.Parse(time.RFC1123, date); err != nil {
			if t, err = time.Parse(time.RFC850, date); err != nil {
				if t, err = time.Parse(time.ANSIC, date); err != nil {
					c.JSON(http.StatusBadRequest, ApiErr{http.StatusBadRequest, "parse HTTP Date error."})
					c.Abort()
					return
				}
			}
		}

		//check client time and server time intervals
		intervals := math.Abs(time.Now().Sub(t).Minutes())
		if intervals > 10 {
			c.JSON(http.StatusUnauthorized, ApiErr{http.StatusUnauthorized, "client time and server time intervals greater than 10 minutes."})
			c.Abort()
			return
		}
		//check Authorization
		//TODO get accesskeyid and secretkey
		accesskeyid := "accesskeyid"
		secretkey := "secretkey"
		if id_key[0] != accesskeyid { //check accesskeyid is exist
			c.JSON(http.StatusUnauthorized, ApiErr{http.StatusUnauthorized, "accesskeyid is not exist."})
			c.Abort()
			return
		}
		c.Request.ParseForm()
		key_value := make([]string, len(c.Request.Form))
		for k, v := range c.Request.Form {
			if len(v) > 0 {
				key_value = append(key_value, k+"="+v[0])
			} else {
				key_value = append(key_value, k+"=")
			}
		}
		sortData := sort.StringSlice(key_value)
		sortData.Sort()
		queryStr := ""
		for _, v := range sortData {
			if queryStr == "" {
				queryStr += "?" + v
			} else {
				queryStr += "&" + v
			}
		}
		stringToSign := c.Request.Method + "\n" +
			date + "\n" +
			c.Request.URL.Path + queryStr
		key := []byte(secretkey)
		mac := hmac.New(sha1.New, key)
		mac.Write([]byte(stringToSign))
		signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		right_authorization := accesskeyid + ":" + signature
		if right_authorization != authorizer {
			c.JSON(http.StatusUnauthorized, ApiErr{http.StatusUnauthorized, "bad authorizer."})
			c.Abort()
			return
		}
		//TODO
		c.Next()
	}
}
