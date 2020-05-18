package vercel

import (
	"encoding/json"
)

func (api VercelAPI) ListRecords(domain string) (*ListRecordsResponse, error) {
	rawResp, err := api.RawListRecords(domain)
	if err != nil {
		return nil, err
	}
	respBody := rawResp.Body()
	if rawResp.StatusCode() == 200 {
		var resp ListRecordsResponse
		json.Unmarshal([]byte(respBody), &resp)
		return &resp, nil
	}
	return nil, api.unmarshalErrorResponse(respBody)
}

func (api VercelAPI) CreateRecord(domain string, type_ string, name string, value string) (*CreateRecordResponse, error) {
	rawResp, err := api.RawCreateRecord(domain, type_, name, value)
	if err != nil {
		return nil, err
	}
	respBody := rawResp.Body()
	if rawResp.StatusCode() == 200 {
		var resp CreateRecordResponse
		json.Unmarshal([]byte(respBody), &resp)
		return &resp, nil
	}
	return nil, api.unmarshalErrorResponse(respBody)
}

func (api VercelAPI) DeleteRecord(domain string, name string) error {
	rawResp, err := api.RawDeleteRecord(domain, name)
	if err != nil {
		return err
	}
	respBody := rawResp.Body()
	if rawResp.StatusCode() == 200 {
		return nil
	}
	return api.unmarshalErrorResponse(respBody)
}
