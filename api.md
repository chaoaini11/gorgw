#GO Rados Gate Way API
	author:		ailn(z.ailn@wmntec.com)
	date:		2016-05-07
	version		1.0
##Version

###获取版本号
	GET /version

	{"Code":200,"Message":"1462610242 api version 1.0"}
##Signature
	signature = accesskeyid + ":" + base64 (
		hmac-sha1 (
			http_verb + "\n" +
			http_date + "\n" +
			path + query
		)
	)

	1.query 生成
	将 url 和 body 中的所有查询参数 格式化为 key=value
	格式化后的 key value 字符串 字典序排序
	用&连接序列化后的 key value 字符串
	如果结果不为空 前边 加 "?"
	eg. ?bucket=mybucket&limit=public
	2.生成待签名字符串
	http_verb + "\n" //http 动词 + 换行
	http_date + "\n" //http Date 请求头(rfc2616 时间格式)
	path + query 	 //URL.Path 根目录 path=/  query
	3.签名编码
	使用用户私钥 对上一步生成的签名字符串 进行 hmac_sha1 签名 并对签名结果进行 base64 编码
	4.构造Signature
	signature = accesskeyid:base64 string	
##Bucket
###创建Bucket
	Request:
		POST /
		Content-Type:application/x-www-form-urlencoded
		Date:{rfc2616 date}
		Authorization:{singature}
	
		bucket={bucketname}&ispublic={false | true}

	Response:
		{Code:200,Message:"create bucket success."}
###列出Bucket
	Request:
		GET /
		Date:{ref2616 date}
		Authorization:{signature}

	Response:
		[
			{
				"Guid":"c00d85cb-6c5f-46fe-a32f-ff972ce38a98",
				"Name":"mybucket",
				"Owner":"cb3a1a34-e1c4-4a49-82dd-88399ec4b138",
				"IsPublic":true
			}
		]
###修改Bucket
	Request:
		PUT /{bucketname}
		Content-Type:application/x-www-form-urlencoded
		Date:{rfc2616 date}
		Authorization:{singature}
	
		ispublic={false | true}

	Response:
		{Code:200,Message:"update bucket success."}

##Object
###创建Object
	Request:(note key must be first param md5 must be second param and file must be third param)
		POST /{bucketname}
		Content-Type:multipart/form-data; boundary=de8ddc1751ba01fa33cbc09a7e61335e27baa2a3e226269e7b0a5cae94cb
		Date:{rfc2616 date}
		Authorization:{signature}
	
		--de8ddc1751ba01fa33cbc09a7e61335e27baa2a3e226269e7b0a5cae94cb
		Content-Disposition: form-data; name="key"
		
		test.txt
		--de8ddc1751ba01fa33cbc09a7e61335e27baa2a3e226269e7b0a5cae94cb
		Content-Disposition: form-data; name="md5"
		
		77645b597f9895345d98217daaa47922
		--de8ddc1751ba01fa33cbc09a7e61335e27baa2a3e226269e7b0a5cae94cb
		Content-Type: application/octet-stream
		Content-Disposition: form-data; name="file"; filename="test.txt"
		
		this is test text
		
		--de8ddc1751ba01fa33cbc09a7e61335e27baa2a3e226269e7b0a5cae94cb--

	Response:
	{Code:200,Message:"create object success."}
###列出Object
	Request:
		GET /{bucketname}
		Date:{rfc2616 date}
		Authorization:{signature}

	Response
		[
			{
				"Guid":"9e501843-e305-4cf8-ad0b-21c91cfab59c",
				"Name":"bucket.go",
				"Bucket":"mybucket",
				"Namespace":"a8e64128-0f64-4740-a0e6-ffd2d234419d",
				"Size":3757,
				"Mime":{
					"Suffix":"",
					"Type":""
				},
				"CreateTime":1463208961,
				"MD5":"86030b1588f8b16c7dbb31700dd68772"
			}
		]

###获取Object
	Request:
		GET /{bucketname}/{objectkey}
		Date:{rfc2616 date}
		Authorization:{signature}

	Reponse
		object data

	
	
