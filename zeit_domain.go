package zeit

import (
	"encoding/json"
)

func (api ZeitAPI) ListDomains() (*ListDomainsResponse, error) {
	rawResp, err := api.RawListDomains()
	if err != nil {
		return nil, err
	}
	respBody := rawResp.Body()
	if rawResp.StatusCode() == 200 {
		var resp ListDomainsResponse
		json.Unmarshal([]byte(respBody), &resp)
		return &resp, nil
	}
	return nil, api.unmarshalErrorResponse(respBody)
}

func (api ZeitAPI) GetDomainPrice(name string) (*DomainPriceResponse, error) {
	rawResp, err := api.RawGetDomainPrice(name)
	if err != nil {
		return nil, err
	}
	respBody := rawResp.Body()
	if rawResp.StatusCode() == 200 {
		var resp DomainPriceResponse
		json.Unmarshal([]byte(respBody), &resp)
		return &resp, nil
	}
	return nil, api.unmarshalErrorResponse(respBody)
}
