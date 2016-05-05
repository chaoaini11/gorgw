/*------------------------
name        config
describe    init api config
author      ailn(z.ailn@wmntec.com)
create      2016-05-05
version     1.0
------------------------*/
package config

import (
	//golang official package
	"fmt"
	"path/filepath"
	"time"

	//third party package
	"github.com/ailncode/golib/config"
)

var Conf map[string]string

func LoadConfig(filePath string) {
	Conf = make(map[string]string)
	file, err := filepath.Abs(filePath)
	if err == nil {
		Conf, err = config.ReadINI(file, "#")
		if err != nil {
			fmt.Println(time.Now().String(), "read ini config error:", err)
		}
	} else {
		fmt.Println(time.Now().String(), "get config file abs file path error:", err)
	}
}
