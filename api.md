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

	POST /
	Date:{rfc2616 date}
	Authorization:{singature}

	bucket={bucketname}&limit={public | private}

	
	
