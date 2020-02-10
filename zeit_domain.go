package zeit

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ListDomainsResponse struct {
	Domains []struct {
		Id                  string   `json:"id"`
		Name                string   `json:"name"`
		ServiceType         string   `json:"serviceType"`
		NsVerifiedAt        int      `json:"nsVerifiedAt"`
		TxtVerifiedAt       int      `json:"txtVerifiedAt"`
		CdnEnabled          bool     `json:"cdnEnabled"`
		CreatedAt           int      `json:"createdAt"`
		BoughtAt            int      `json:"boughtAt"`
		VerificationRecord  string   `json:"verificationRecord"`
		Verified            bool     `json:"verified"`
		Nameservers         []string `json:"nameservers"`
		IntendedNameservers []string `json:"intendedNameservers"`
		Creator             struct {
			Id       string `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		} `json:"creator"`
	} `json:"domains"`
}

func (api ZeitAPI) RawListDomains() (*resty.Response, error) {
	return api.get("/v4/domains", map[string]string{})
}

func (api ZeitAPI) ListDomains() (*ListDomainsResponse, error) {
	rawResp, err := api.RawListDomains()
	if err != nil {
		return nil, err
	}
	respBody := rawResp.Body()
	var resp ListDomainsResponse
	json.Unmarshal([]byte(respBody), &resp)
	fmt.Println(resp)
	return &resp, nil
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
