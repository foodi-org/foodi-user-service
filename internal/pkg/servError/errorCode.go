package servError

/*
用户中心码段: 110000 ~ 119999

* `账号服务`: 111000 ~ 111999
* `用户服务`: 112000 ~ 112999
* `文章服务`: 113000 ~ 113999
*/

const (
	Common  = 111000
	UNKNOWN = 111099

	InvalidParam  = 113000 // 无效的参数
	InvalidAction = 113001 // 无效的操作类型

	ArticleNotFound = 113010 // 文章不存在
)
