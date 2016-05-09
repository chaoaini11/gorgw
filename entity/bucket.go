/*------------------------
name        bucket
describe    bucket entity
author      ailn(z.ailn@wmntec.com)
create      2016-05-07
version     1.0
------------------------*/
package entity

type Bucket struct {
	Guid     string //ceph namespace guid
	Name     string //user bucket name
	Owner    string //bucket owner guid
	IsPublic bool
}
