package zeit

import "github.com/go-resty/resty/v2"

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
