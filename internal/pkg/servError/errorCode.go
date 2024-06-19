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

	MissPhoneNumber    = 111000
	MissPassword       = 111001
	LoginCodeError     = 111002
	LoginPasswordError = 111003
	LoginFailError     = 111004

	PhoneNumberNotFound = 111010 // 电话号码不存在

	AccountNotBindPassword = 111011 // 账号没有绑定密码

	AccountPasswordNotMatch = 111012

	InvalidUserCoup = 111013

	InvalidRegisterCoup = 111014

	ArticleNotFound = 113010 // 文章不存在

	ArticleCreateFail = 113011 // 创建文章失败

	ArticleSaveFail = 113012 // 文章收藏失败

	CommentNotFound = 113020
)
