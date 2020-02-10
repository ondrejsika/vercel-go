package zeit

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ZeitAPI struct {
	token     string
	apiOrigin string
}

func New(token string) ZeitAPI {
	api := ZeitAPI{token, "https://api.zeit.co"}
	return api
}

func NewOrigin(token string, apiOrigin string) ZeitAPI {
	api := ZeitAPI{token, apiOrigin}
	return api
}

func (api ZeitAPI) get(url string, query map[string]string) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetAuthToken(api.token).SetQueryParams(query).Get(api.apiOrigin + url)
}

func (api ZeitAPI) post(url string, body map[string]interface{}) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetAuthToken(api.token).SetBody(body).Post(api.apiOrigin + url)
}

func (api ZeitAPI) delete(url string) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetAuthToken(api.token).Delete(api.apiOrigin + url)
}

type errorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (api ZeitAPI) unmarshalErrorResponse(respBody []byte) error {
	var resp errorResponse
	json.Unmarshal([]byte(respBody), &resp)
	return fmt.Errorf("Zeit API Error: %s: %s", resp.Error.Code, resp.Error.Message)
}
