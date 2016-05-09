package main

import (
	"fmt"
	"github.com/ailncode/golib/mongo"
	. "github.com/ailncode/gorgw/config"
	"github.com/ailncode/gorgw/entity"
	"github.com/pborman/uuid"
	"os"
)

func main() {
	if len(os.Args) == 6 {
		LoadConfig("../config/gorgw.conf")
		mgo, err := mongo.NewMongo(Conf["server"])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer mgo.Close()
		err = mgo.Insert(Conf["db"], Conf["usercoll"], &entity.User{uuid.New(), os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Args[5]})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("create user success.")
	} else {
		fmt.Println("Useage:")
		fmt.Println("createuser [account] [pwd] [name] [accesskeyid] [secretkey]")
	}
}
