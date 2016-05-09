package main

import (
	"fmt"
	"github.com/ailncode/golib/base64"
	"github.com/ailncode/golib/hmacsha1"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	accesskeyid = "01"
	secretkey   = "abcdefg"
	apiurl      = "http://127.0.0.1:8612/"
)

func list() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Now().Format(time.RFC1123)
	str := req.Method + "\n" +
		t + "\n" +
		"/"
	sign := base64.Encode(hmacsha1.Encrypt(secretkey, []byte(str)))
	sign = accesskeyid + ":" + sign
	req.Header.Add("Date", t)
	req.Header.Add("Authorization", sign)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func create(name string, ispublic string) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiurl, strings.NewReader("bucketname="+name+"&ispublic="+ispublic))
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Now().Format(time.RFC1123)
	str := req.Method + "\n" +
		t + "\n" +
		"/?bucketname=" + name + "&ispublic=" + ispublic
	sign := base64.Encode(hmacsha1.Encrypt(secretkey, []byte(str)))
	sign = accesskeyid + ":" + sign
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Date", t)
	req.Header.Add("Authorization", sign)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func update(name string, ispublic string) {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", apiurl+name, strings.NewReader("ispublic="+ispublic))
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Now().Format(time.RFC1123)
	str := req.Method + "\n" +
		t + "\n" +
		"/" + name + "?ispublic=" + ispublic
	sign := base64.Encode(hmacsha1.Encrypt(secretkey, []byte(str)))
	sign = accesskeyid + ":" + sign
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Date", t)
	req.Header.Add("Authorization", sign)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func main() {
	if len(os.Args) == 1 {
		fmt.Println("bucket version 1.0")
		fmt.Println("use bucket help for help")
		return
	}
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "help":
			fmt.Println("bucket help")
			return
		case "list":
			list()
			return
		default:
			fmt.Println("use bucket help for help")
			return
		}
	}
	if len(os.Args) > 2 {
		switch os.Args[1] {
		case "create":
			if len(os.Args) == 3 {
				create(os.Args[2], "false")
				return
			}
			if len(os.Args) == 4 {
				if os.Args[3] == "true" || os.Args[3] == "false" {
					create(os.Args[2], os.Args[3])
					return
				}
			}
			useCreate()
		case "update":
			if len(os.Args) == 3 {
				update(os.Args[2], "false")
				return
			}
			if len(os.Args) == 4 {
				if os.Args[3] == "true" || os.Args[3] == "false" {
					update(os.Args[2], os.Args[3])
					return
				}
			}
			useUpdate()
		default:
			use()
			return
		}
	}
}
func use() {
}
func help() {
	fmt.Println("use bucket help for help")
}
func useCreate() {
	fmt.Println("bucket create [bucketname] [ispublic true | false]")
}
func useUpdate() {
	fmt.Println("bucket update [bucketname] [ispublic true | false]")
}
