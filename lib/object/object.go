/*------------------------
name        object
describe    object lib
author      ailn(z.ailn@wmntec.com)
create      2016-05-11
version     1.0
------------------------*/
package object

import (
	//golang official package
	"bufio"
	"fmt"
	"mime/multipart"

	//third party package
	"github.com/ailncode/gorados"

	//project package
	. "github.com/ailncode/gorgw/config"
)

func Create(nameSpace, key string, file multipart.File, taskId string) {
	//conf
	defer file.Close()
	conf := &gorados.Config{Conf["cephcluster"], Conf["cephuser"], Conf["cephpool"], Conf["ceph"], nameSpace}
	//new Rados
	r, err := gorados.New(conf)
	if err != nil {
		//write task result
		return
	}
	//close
	defer r.Close()
	b := r.NewBuffer(key)
	defer b.Close()
	buf := bufio.NewReader(file)
	write_len, err := buf.WriteTo(b)
	if err != nil {
		//write task result
		fmt.Println(err)
		return
	}
	fmt.Println("write success.", write_len)
	//write task result
}
