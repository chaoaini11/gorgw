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
	"encoding/hex"
	"fmt"
	"io"

	//third party package
	"github.com/ailncode/gorados"

	//project package
	. "github.com/ailncode/gorgw/config"
)

func Create(nameSpace, key, md5 string, rc io.ReadCloser, taskId string) {
	//conf
	defer rc.Close()
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
	buf := bufio.NewReader(rc)
	write_len, err := buf.WriteTo(b)
	if err != nil {
		//write task result
		fmt.Println(err)
		return
	}
	//check sum
	md5_sum := b.MD5.Sum(nil)
	if md5 != hex.EncodeToString(md5_sum) {
		//wirte error
		fmt.Println(md5, hex.EncodeToString(md5_sum))
	}
	fmt.Println("write success.", write_len)
	//write task result
}
