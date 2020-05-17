package vercel

import "github.com/go-resty/resty/v2"

func (api VercelAPI) RawListRecords(domain string, type_ string, name string, value string) (*resty.Response, error) {
	return api.get("/v2/domains/"+domain+"/records", map[string]string{})
}

func (api VercelAPI) RawCreateRecord(domain string, type_ string, name string, value string) (*resty.Response, error) {
	return api.post("/v2/domains/"+domain+"/records", map[string]interface{}{
		"type":  type_,
		"name":  name,
		"value": value,
	})
}

func (api VercelAPI) RawDeleteRecord(domain string, recId string) (*resty.Response, error) {
	return api.delete("/v2/domains/" + domain + "/records/" + recId)
}
