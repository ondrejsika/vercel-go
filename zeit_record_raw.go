package zeit

import "github.com/go-resty/resty/v2"

func (api ZeitAPI) RawListRecords(domain string, type_ string, name string, value string) (*resty.Response, error) {
	return api.get("/v2/domains/"+domain+"/records", map[string]string{})
}

func (api ZeitAPI) RawCreateRecord(domain string, type_ string, name string, value string) (*resty.Response, error) {
	return api.post("/v2/domains/"+domain+"/records", map[string]interface{}{
		"type":  type_,
		"name":  name,
		"value": value,
	})
}

func (api ZeitAPI) RawDeleteRecord(domain string, recId string) (*resty.Response, error) {
	return api.delete("/v2/domains/" + domain + "/records/" + recId)
}