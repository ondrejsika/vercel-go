package zeit

import "github.com/go-resty/resty/v2"

func (api ZeitAPI) DomainBuy(name string, expectedPrice int) (*resty.Response, error) {
	return api.post("/v2/domains/buy", map[string]interface{}{
		"name":          name,
		"expectedPrice": expectedPrice,
	})
}

func (api ZeitAPI) DomainPrice(name string) (*resty.Response, error) {
	return api.get("/v4/domains/price", map[string]string{
		"name": name,
	})
}
