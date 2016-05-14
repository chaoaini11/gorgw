package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/ailncode/golib/base64"
	"github.com/ailncode/golib/hmacsha1"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

var (
	accesskeyid = "01"
	secretkey   = "abcdefg"
	apiurl      = "http://192.168.99.156:8610/"
)

func create(key, filename string) {
	fmt.Println(key)
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	keyKey, err := w.CreateFormField("key")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = keyKey.Write([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	m := md5.New()
	io.Copy(m, f)

	md5Str := hex.EncodeToString(m.Sum(nil))
	md5Key, err := w.CreateFormField("md5")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = md5Key.Write([]byte(md5Str))
	if err != nil {
		fmt.Println(err)
		return
	}
	fileKey, err := w.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Seek(0, 0)
	cop_len, err := io.Copy(fileKey, f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cop_len)
	w.Close()
	client := &http.Client{}
	req, err := http.NewRequest("POST", apiurl+"mybucket", &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Now().Format(time.RFC1123)
	str := req.Method + "\n" +
		t + "\n" +
		"/mybucket"
	sign := base64.Encode(hmacsha1.Encrypt(secretkey, []byte(str)))
	sign = accesskeyid + ":" + sign
	fmt.Println(sign)
	//req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Content-Type", w.FormDataContentType())
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
func list(bucketName string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiurl+bucketName, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	t := time.Now().Format(time.RFC1123)
	str := req.Method + "\n" +
		t + "\n" +
		"/" + bucketName
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
func main() {
	if len(os.Args) == 1 {
		fmt.Println("bucket version 1.0")
		fmt.Println("use bucket help for help")
		return
	}
	if len(os.Args) == 2 {
	}
	if len(os.Args) > 2 {
		switch os.Args[1] {
		case "create":
			if len(os.Args) == 4 {
				create(os.Args[2], os.Args[3])
			} else {
				useCreate()
			}
			break
		case "list":
			if len(os.Args) == 3 {
				list(os.Args[2])
			} else {
				useList()
			}
		}
	}
}
func useCreate() {
	fmt.Println("object create [key] [path]")
}
func useList() {
	fmt.Println("object list [bucket]")
}
