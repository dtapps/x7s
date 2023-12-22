package x7s

func (c *Client) GetApiURL() string {
	return c.config.apiURL
}

func (c *Client) GetPartnerID() int64 {
	return c.config.partnerID
}

func (c *Client) GetApiKey() string {
	return c.config.apiKey
}
