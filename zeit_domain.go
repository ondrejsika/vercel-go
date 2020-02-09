package zeit

import "github.com/go-resty/resty/v2"

func (api ZeitAPI) RawListDomains() (*resty.Response, error) {
	return api.get("/v4/domains", map[string]string{})
}

func (api ZeitAPI) RawAddDomain(name string) (*resty.Response, error) {
	return api.post("/v4/domains", map[string]interface{}{
		"name": name,
	})
}

func (api ZeitAPI) RawTransferDomainIn(name string, expectedPrice int, authCode string) (*resty.Response, error) {
	return api.post("/v4/domains", map[string]interface{}{
		"name":          name,
		"method":        "transfer-in",
		"authCode":      authCode,
		"expectedPrice": expectedPrice,
	})
}

func (api ZeitAPI) RawVerifyDomain(name string) (*resty.Response, error) {
	return api.post("/v4/domains/"+name+"/verify", map[string]interface{}{})
}

func (api ZeitAPI) RawGetDomain(name string) (*resty.Response, error) {
	return api.get("/v4/domains/"+name, map[string]string{})
}

func (api ZeitAPI) RawRemoveDomain(domain string) (*resty.Response, error) {
	return api.delete("/v4/domains/" + domain)
}

func (api ZeitAPI) RawCheckDomainAvailibility(name string) (*resty.Response, error) {
	return api.get("/v4/domains/status", map[string]string{
		"name": name,
	})
}

func (api ZeitAPI) RawBuyDomain(name string, expectedPrice int) (*resty.Response, error) {
	return api.post("/v4/domains/buy", map[string]interface{}{
		"name":          name,
		"expectedPrice": expectedPrice,
	})
}

func (api ZeitAPI) RawGetDomainPrice(name string) (*resty.Response, error) {
	return api.get("/v4/domains/price", map[string]string{
		"name": name,
	})
}
