package vercel

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type VercelAPI struct {
	token     string
	apiOrigin string
}

func New(token string) VercelAPI {
	api := VercelAPI{token, "https://api.vercel.com"}
	return api
}

func NewOrigin(token string, apiOrigin string) VercelAPI {
	api := VercelAPI{token, apiOrigin}
	return api
}

func (api VercelAPI) get(url string, query map[string]string) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetAuthToken(api.token).SetQueryParams(query).Get(api.apiOrigin + url)
}

func (api VercelAPI) post(url string, body map[string]interface{}) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetAuthToken(api.token).SetBody(body).Post(api.apiOrigin + url)
}

func (api VercelAPI) delete(url string) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetAuthToken(api.token).Delete(api.apiOrigin + url)
}

type errorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (api VercelAPI) unmarshalErrorResponse(respBody []byte) error {
	var resp errorResponse
	json.Unmarshal([]byte(respBody), &resp)
	return fmt.Errorf("Vercel API Error: %s: %s", resp.Error.Code, resp.Error.Message)
}
