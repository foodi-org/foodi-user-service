package enumgen

type (
	// CardType 证件类型
	CardType int8

	// GenderType 性别
	GenderType string

	// UserType 用户类型
	UserType string

	// ArticleType 文章类型
	ArticleType string
)

const (
	FirstGenerationIDCard  CardType = 1 // 第一代居民身份证
	SecondGenerationIDCard CardType = 2 // 第二代居民身份证
	BusinessLicense        CardType = 3 // 营业执照
)

const (
	MALE   GenderType = "male"
	FEMALE GenderType = "female"
)

const (
	NormalUser   UserType = "consumer" // 普通用户
	MerchantUser UserType = "merchant" // 商家用户
)

const (
	Video ArticleType = "video"
	Text  ArticleType = "text"
)
