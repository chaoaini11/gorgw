/*------------------------
name		base
describe	base library
author		ailn(z.ailn@wmntec.com)
create		2016-05-05
version		1.0
------------------------*/
package base

import (
	//golang official package
	"encoding/json"
)

//api custom error
type ApiErr struct {
	Code    int
	Message string
}

//method of ApiErr to achieve interface error
func (apiErr *ApiErr) Error() string {
	b, err := json.Marshal(apiErr)
	if err != nil {
		return `{"Code":500,"Message":"Serialize api error error."}`
	}
	return string(b)
}
