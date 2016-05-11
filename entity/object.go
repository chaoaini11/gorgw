/*------------------------
name        object
describe    object entity
author      ailn(z.ailn@wmntec.com)
create      2016-05-08
version     1.0
------------------------*/
package entity

import (
	//third party package
	"github.com/ailncode/golib/mime"
)

type Object struct {
	Guid       string //guid
	Name       string //the object name
	Bucket     string //which bucket this object in
	Namespace  string //which namespace this object in
	Size       int64  //Byte
	Mime       mime.Mime
	CreateTime int64
}
