package x7s

func (c *Client) SetApiURL(apiURL string) {
	c.config.apiURL = apiURL
}

func (c *Client) SetPartnerID(partnerID int64) {
	c.config.partnerID = partnerID
}

func (c *Client) SetApiKey(apiKey string) {
	c.config.apiKey = apiKey
}
