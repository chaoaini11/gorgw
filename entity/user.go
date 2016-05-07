/*------------------------
name        user
describe    user entity
author      ailn(z.ailn@wmntec.com)
create      2016-05-07
version     1.0
------------------------*/
package entity

type User struct {
	Account     string
	Password    string
	Name        string
	AccessKeyID string
	SecretKey   string
}
