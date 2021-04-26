package common

// Token
type AccessToken struct {
	// 授权SessionKey
	AccessToken string
	// 第三方平台Token
	ThirdPlatToken string
	// SessionKey过期时间间隔(单位：秒)
	ExpireIn int

	//SessionKey过期时间
	ExpireTime string
	TokenType  string
	// 刷新SessionKey(当过期时)所需RefreshToken
	RefreshToken string
	//RefreshToken过期时间
	RefreshTokenExpireTime string
	// 店铺ID
	VenderId string `json:"venderid"`
	//用户昵称
	Nickname string `json:"nickname"`
	//店铺类别（B代表天猫店铺，C代表 集市卖家
	Subplat string   `json:"subplat"`
	Scope   []string `json:"scope"`
}
