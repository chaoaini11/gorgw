/*------------------------
name        bucket
describe    bucket lib
author      ailn(z.ailn@wmntec.com)
create      2016-05-11
version     1.0
------------------------*/
package bucket

import (
	//golang official package
	"errors"
	"fmt"

	//third party package
	"github.com/ailncode/golib/mongo"

	//project package
	. "github.com/ailncode/gorgw/config"
	"github.com/ailncode/gorgw/entity"
)

func Get(owner, bucketName string) (entity.Bucket, error) {
	bucket := entity.Bucket{}
	mgo, err := mongo.NewMongo(Conf["server"])
	if err != nil {
		return bucket, err
	}
	defer mgo.Close()
	err = mgo.FindOne(Conf["db"], Conf["bucketcoll"], map[string]interface{}{"owner": owner, "name": bucketName}, &bucket)
	return bucket, nil
}
func List(owner, bucketName string) ([]entity.Object, error) {
	var objectList []entity.Object
	b, err := Get(owner, bucketName)
	if err != nil {
		return objectList, err
	}
	mgo, err := mongo.NewMongo(Conf["server"])
	if err != nil {
		return objectList, err
	}
	defer mgo.Close()
	err = mgo.FindAll(Conf["db"], Conf["objectcoll"], map[string]interface{}{"namespace": b.Guid}, &objectList)
	return objectList, err
}
func IsExist(owner, bucketName string) (bool, error) {
	mgo, err := mongo.NewMongo(Conf["server"])
	if err != nil {
		fmt.Println(err)
		return false, errors.New("open mongodb server error.")
	}
	defer mgo.Close()
	return mgo.IsExist(Conf["db"], Conf["bucketcoll"], map[string]interface{}{"name": bucketName, "owner": owner}), nil
}
