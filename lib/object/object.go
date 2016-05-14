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
	"time"

	//third party package
	"github.com/ailncode/golib/mime"
	"github.com/ailncode/golib/mongo"
	"github.com/ailncode/gorados"
	"github.com/pborman/uuid"

	//project package
	. "github.com/ailncode/gorgw/config"
	"github.com/ailncode/gorgw/entity"
	"github.com/ailncode/gorgw/lib/bucket"
)

func Get(owner, bucketName, key string) (entity.Object, error) {
	var o entity.Object
	b, err := bucket.Get(owner, bucketName)
	if err != nil {
		return o, err
	}
	mgo, err := mongo.NewMongo(Conf["server"])
	if err != nil {
		fmt.Println(err)
		return o, errors.New("open mongodb server error.")
	}
	defer mgo.Close()
	err = mgo.FindOne(Conf["db"], Conf["objectcoll"],
		map[string]interface{}{"namespace": b.Guid, "name": key}, &o)
	return o, err
}
func GetReader(owner, bucketName, key string) (*gorados.RadosBuffer, error) {

	var buf *gorados.RadosBuffer
	obj, err := Get(owner, bucketName, key)
	if err != nil {
		return buf, err
	}
	conf := &gorados.Config{Conf["cephcluster"], Conf["cephuser"], Conf["cephpool"], Conf["cephconfig"], obj.Namespace}
	//new Rados
	r, err := gorados.New(conf)
	if err != nil {
		fmt.Println(err)
		return buf, err
	}
	return r.NewBuffer(obj.Guid), nil
}
func Create(nameSpace, key, bucketName, md5 string, rc io.ReadCloser, taskId string) error {
	//conf
	defer rc.Close()
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
	object_guid := uuid.New()
	b := r.NewBuffer(object_guid)
	defer b.Close()
	buf := bufio.NewReaderSize(rc, 1024*1024*4)
	_, err = buf.WriteTo(b)
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
	mgo, err := mongo.NewMongo(Conf["server"])
	if err != nil {
		fmt.Println(err)
		return errors.New("open mongodb server error.")
	}
	m := mime.Suffix(key)
	if m == mime.UKNOWN {
		m = mime.FileHeader(b.FileHeader)
	}
	obj := entity.Object{object_guid, key, bucketName, nameSpace,
		b.Off, m, time.Now().Unix(), md5}
	defer mgo.Close()
	err = mgo.Insert(Conf["db"], Conf["objectcoll"], &obj)
	if err != nil {
		return err
	}
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
