package zeit

import "github.com/go-resty/resty/v2"

func (api ZeitAPI) RecordCreate(domain string, type_ string, name string, value string) (*resty.Response, error) {
	return api.post("/v2/domains/"+domain+"/records", map[string]interface{}{
		"type":  type_,
		"name":  name,
		"value": value,
	})
}
