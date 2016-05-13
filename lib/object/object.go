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
	"errors"
	"fmt"
	"io"

	//third party package
	"github.com/ailncode/golib/mongo"
	"github.com/ailncode/gorados"

	//project package
	. "github.com/ailncode/gorgw/config"
)

func Create(nameSpace, key, md5 string, rc io.ReadCloser, taskId string) error {
	//conf
	defer rc.Close()
	fmt.Println(rc)
	fmt.Println("create config")
	conf := &gorados.Config{Conf["cephcluster"], Conf["cephuser"], Conf["cephpool"], Conf["cephconfig"], nameSpace}
	//new Rados
	r, err := gorados.New(conf)
	if err != nil {
		fmt.Println(err)
		//write task result
		return err
	}
	//close
	defer r.Close()
	fmt.Println("create new buffer")
	b := r.NewBuffer(key)
	defer b.Close()
	buf := bufio.NewReaderSize(rc, 1024*1024*64)
	write_len, err := buf.WriteTo(b)
	if err != nil {
		//write task result
		fmt.Println(err)
		return err
	}
	//check sum
	md5_sum := b.MD5.Sum(nil)
	if md5 != hex.EncodeToString(md5_sum) {
		return errors.New("check md5 failed.")
		//wirte error
		fmt.Println(md5, hex.EncodeToString(md5_sum))
	}
	fmt.Println("write success.", write_len)
	return nil
	//write task result
}

func IsExist(bucketGuid, objectKey string) (bool, error) {
	mgo, err := mongo.NewMongo(Conf["server"])
	if err != nil {
		fmt.Println(err)
		return false, errors.New("open mongodb server error.")
	}
	defer mgo.Close()
	return mgo.IsExist(Conf["db"], Conf["objectcoll"], map[string]interface{}{"namespace": bucketGuid, "name": objectKey}), nil
}
