package config

// QQ qq 登录配置，详情请见 https://connect.qq.com/
type QQ struct {
	Enable      bool   `json:"enable" yaml:"enable"`             // 是否启用 qq 登录，true 表示启用，false 表示禁用
	AppID       string `json:"app_id" yaml:"app_id"`             // 应用 ID
	AppKey      string `json:"app_key" yaml:"app_key"`           // 应用密钥
	RedirectURI string `json:"redirect_uri" yaml:"redirect_uri"` // 网站回调域
}
