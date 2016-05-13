/*------------------------
name        task
describe    task entity
author      ailn(z.ailn@wmntec.com)
create      2016-05-11
version     1.0
------------------------*/
package entity

import (
//golang official package
)

type Task struct {
	Guid       string //guid
	CreateTime int64
	Describe   string
	Status     uint
	Message    string
}
