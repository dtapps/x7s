package x7s

import (
	"errors"
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL    string `json:"api_url"`    // 接口地址
	PartnerID int64  `json:"partner_id"` // 平台分配商户号
	ApiKey    string `json:"api_key"`    // 渠道分配的密钥
}

// Client 实例
type Client struct {
	config struct {
		apiURL    string // 接口地址
		partnerID int64  // 平台分配商户号
		apiKey    string // 渠道分配的密钥
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.apiURL = config.ApiURL
	c.config.partnerID = config.PartnerID
	c.config.apiKey = config.ApiKey

	if c.config.apiURL == "" {
		return nil, errors.New("需要配置ApiURL")
	}

	return c, nil
}
