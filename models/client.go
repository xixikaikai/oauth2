package models

// NewClient 创建客户端模型实例
func NewClient() *Client {
	return &Client{}
}

// Client 客户端信息
type Client struct {
	ID     string // 客户端ID
	Secret string // 密钥
	Domain string // 域名url
}

// GetID 客户端ID
func (c *Client) GetID() string {
	return c.ID
}

// GetSecret 客户端秘钥
func (c *Client) GetSecret() string {
	return c.Secret
}

// GetDomain 域名URL
func (c *Client) GetDomain() string {
	return c.Domain
}

// GetExtraData 扩展数据
func (c *Client) GetExtraData() interface{} {
	return nil
}
