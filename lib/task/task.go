/*------------------------
name        task
describe    task lib
author      ailn(z.ailn@wmntec.com)
create      2016-05-13
version     1.0
------------------------*/
package task

import (
	//third party package
	"github.com/ailncode/golib/mongo"

	//project package
	. "github.com/ailncode/gorgw/config"
	"github.com/ailncode/gorgw/entity"
)

func Create(task *entity.Task) error {
	mgo, err := mongo.NewMongo(Conf["server"])
	if err != nil {
		return err
	}
	defer mgo.Close()
	err = mgo.Insert(Conf["db"], Conf["taskcoll"], task)
	return err
}
